package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
)

const (
	defaultFilePath = "data/all-stock.db"
	defaultRegion   = "auto"
)

type uploadConfig struct {
	FilePath        string
	Bucket          string
	ObjectKey       string
	AccountID       string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	ContentType     string
	Timeout         time.Duration
	Verify          bool
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Upload file to Cloudflare R2.\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Env fallbacks:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_BUCKET\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_OBJECT_KEY\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_ACCOUNT_ID\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_ENDPOINT\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_ACCESS_KEY_ID or AWS_ACCESS_KEY_ID\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  R2_SECRET_ACCESS_KEY or AWS_SECRET_ACCESS_KEY\n\n")
		flag.PrintDefaults()
	}

	cfg := uploadConfig{}
	flag.StringVar(&cfg.FilePath, "file", envFirst("R2_FILE_PATH", defaultFilePath), "Local file path to upload")
	flag.StringVar(&cfg.Bucket, "bucket", envFirst("R2_BUCKET", ""), "Cloudflare R2 bucket name")
	flag.StringVar(&cfg.ObjectKey, "key", envFirst("R2_OBJECT_KEY", ""), "Destination object key in bucket, default is file base name")
	flag.StringVar(&cfg.AccountID, "account-id", envFirst("R2_ACCOUNT_ID", ""), "Cloudflare account ID, used to build R2 endpoint when -endpoint is empty")
	flag.StringVar(&cfg.Endpoint, "endpoint", envFirst("R2_ENDPOINT", ""), "Full R2 S3 endpoint, e.g. https://<accountid>.r2.cloudflarestorage.com")
	flag.StringVar(&cfg.AccessKeyID, "access-key-id", envFirst("R2_ACCESS_KEY_ID", envFirst("AWS_ACCESS_KEY_ID", "")), "R2 access key ID")
	flag.StringVar(&cfg.SecretAccessKey, "secret-access-key", envFirst("R2_SECRET_ACCESS_KEY", envFirst("AWS_SECRET_ACCESS_KEY", "")), "R2 secret access key")
	flag.StringVar(&cfg.ContentType, "content-type", "application/octet-stream", "Content-Type sent to R2")
	flag.DurationVar(&cfg.Timeout, "timeout", 30*time.Minute, "Upload timeout")
	flag.BoolVar(&cfg.Verify, "verify", true, "Run HEAD after upload and compare remote size")
	flag.Parse()

	if err := cfg.normalize(); err != nil {
		logrus.Fatalf("invalid config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := uploadFile(ctx, cfg); err != nil {
		logrus.Fatalf("upload failed: %v", err)
	}
	logrus.Infof("upload finished: file=%s bucket=%s key=%s", cfg.FilePath, cfg.Bucket, cfg.ObjectKey)
}

func (cfg *uploadConfig) normalize() error {
	if cfg.FilePath == "" {
		cfg.FilePath = defaultFilePath
	}
	if cfg.ObjectKey == "" {
		cfg.ObjectKey = filepath.Base(cfg.FilePath)
	}
	if cfg.Endpoint == "" {
		if cfg.AccountID == "" {
			return fmt.Errorf("missing R2 endpoint: set -endpoint or -account-id")
		}
		cfg.Endpoint = fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountID)
	}
	missing := make([]string, 0, 3)
	if cfg.Bucket == "" {
		missing = append(missing, "bucket")
	}
	if cfg.AccessKeyID == "" {
		missing = append(missing, "access-key-id")
	}
	if cfg.SecretAccessKey == "" {
		missing = append(missing, "secret-access-key")
	}
	if len(missing) != 0 {
		return fmt.Errorf("missing required config: %s", strings.Join(missing, ", "))
	}
	if cfg.Timeout <= 0 {
		return fmt.Errorf("timeout must be positive")
	}
	return nil
}

func uploadFile(ctx context.Context, cfg uploadConfig) error {
	info, err := os.Stat(cfg.FilePath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("file path is directory: %s", cfg.FilePath)
	}

	checksum, err := sha256File(cfg.FilePath)
	if err != nil {
		return fmt.Errorf("compute sha256: %w", err)
	}

	file, err := os.Open(cfg.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	awsCfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(defaultRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		return fmt.Errorf("load aws config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(options *s3.Options) {
		options.BaseEndpoint = aws.String(cfg.Endpoint)
		options.UsePathStyle = true
	})

	uploadCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	putOutput, err := client.PutObject(uploadCtx, &s3.PutObjectInput{
		Bucket:        aws.String(cfg.Bucket),
		Key:           aws.String(cfg.ObjectKey),
		Body:          file,
		ContentLength: aws.Int64(info.Size()),
		ContentType:   aws.String(cfg.ContentType),
	})
	if err != nil {
		return fmt.Errorf("put object: %w", err)
	}

	if cfg.Verify {
		headOutput, err := client.HeadObject(uploadCtx, &s3.HeadObjectInput{
			Bucket: aws.String(cfg.Bucket),
			Key:    aws.String(cfg.ObjectKey),
		})
		if err != nil {
			return fmt.Errorf("head object: %w", err)
		}
		if headOutput.ContentLength == nil {
			return fmt.Errorf("head object returned empty content length")
		}
		if aws.ToInt64(headOutput.ContentLength) != info.Size() {
			return fmt.Errorf("remote size mismatch: local=%d remote=%d", info.Size(), aws.ToInt64(headOutput.ContentLength))
		}
	}

	logrus.Infof("uploaded: bucket=%s key=%s bytes=%d sha256=%s etag=%s endpoint=%s", cfg.Bucket, cfg.ObjectKey, info.Size(), checksum, aws.ToString(putOutput.ETag), cfg.Endpoint)
	return nil
}

func sha256File(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func envFirst(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}