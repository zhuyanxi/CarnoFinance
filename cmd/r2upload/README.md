# r2upload

`r2upload` 用 S3 兼容接口将本地文件上传到 Cloudflare R2。

默认上传文件是 `data/all-stock.db`。

## 功能

- 上传本地文件到指定 R2 bucket
- 默认对象 key 使用文件名 basename
- 若只提供 `account-id`，自动拼出 R2 endpoint
- 上传前计算文件 SHA256，并在日志中输出
- 默认上传后执行 `HEAD` 校验远端文件大小

## 前置条件

- 已创建 Cloudflare R2 bucket
- 已创建 R2 Access Key / Secret Key
- 本地待上传文件存在

## 必填配置

下面配置必须提供：

- `bucket`
- `access-key-id`
- `secret-access-key`
- `account-id` 或 `endpoint` 二选一

## 快速开始

先设置环境变量：

```bash
export R2_BUCKET=my-backup-bucket
export R2_ACCOUNT_ID=your-cloudflare-account-id
export R2_ACCESS_KEY_ID=your-r2-access-key-id
export R2_SECRET_ACCESS_KEY=your-r2-secret-access-key
```

上传默认文件 `data/all-stock.db`：

```bash
go run ./cmd/r2upload -key backups/all-stock.db
```

也可用 Makefile：

```bash
make upload-r2
```

若不传 `-key`，默认对象 key 是 `all-stock.db`。

## 环境变量

| 环境变量 | 是否必填 | 说明 |
| --- | --- | --- |
| `R2_FILE_PATH` | 否 | 本地文件路径，默认 `data/all-stock.db` |
| `R2_BUCKET` | 是 | R2 bucket 名 |
| `R2_OBJECT_KEY` | 否 | 远端对象 key，默认取文件 basename |
| `R2_ACCOUNT_ID` | 条件必填 | Cloudflare account id。未传 `R2_ENDPOINT` 时必填 |
| `R2_ENDPOINT` | 条件必填 | 完整 endpoint，例如 `https://<account-id>.r2.cloudflarestorage.com` |
| `R2_ACCESS_KEY_ID` | 是 | R2 access key id |
| `R2_SECRET_ACCESS_KEY` | 是 | R2 secret access key |
| `AWS_ACCESS_KEY_ID` | 否 | `R2_ACCESS_KEY_ID` 备用值 |
| `AWS_SECRET_ACCESS_KEY` | 否 | `R2_SECRET_ACCESS_KEY` 备用值 |

## 命令行参数

| 参数 | 默认值 | 说明 |
| --- | --- | --- |
| `-file` | `data/all-stock.db` | 本地文件路径 |
| `-bucket` | 空 | R2 bucket 名 |
| `-key` | 文件 basename | 远端对象 key |
| `-account-id` | 空 | 自动生成 endpoint 用 |
| `-endpoint` | 空 | 直接指定完整 R2 endpoint |
| `-access-key-id` | 空 | R2 access key id |
| `-secret-access-key` | 空 | R2 secret access key |
| `-content-type` | `application/octet-stream` | 上传时 Content-Type |
| `-timeout` | `30m` | 上传超时 |
| `-verify` | `true` | 上传后是否用 `HEAD` 校验大小 |

## 示例

上传默认 DB，key 放到备份目录：

```bash
go run ./cmd/r2upload \
  -bucket my-backup-bucket \
  -account-id 1234567890abcdef \
  -access-key-id xxx \
  -secret-access-key yyy \
  -key backups/all-stock.db
```

上传自定义文件：

```bash
go run ./cmd/r2upload \
  -file data/another.db \
  -bucket my-backup-bucket \
  -account-id 1234567890abcdef \
  -access-key-id xxx \
  -secret-access-key yyy \
  -key backups/another.db
```

直接指定完整 endpoint：

```bash
go run ./cmd/r2upload \
  -bucket my-backup-bucket \
  -endpoint https://1234567890abcdef.r2.cloudflarestorage.com \
  -access-key-id xxx \
  -secret-access-key yyy
```

跳过上传后校验：

```bash
go run ./cmd/r2upload \
  -bucket my-backup-bucket \
  -account-id 1234567890abcdef \
  -access-key-id xxx \
  -secret-access-key yyy \
  -verify=false
```

## 日志输出

上传成功后，日志会输出这些信息：

- bucket
- key
- bytes
- sha256
- etag
- endpoint

## 注意

- 当前实现使用单次 `PutObject` 上传，不是 multipart upload
- 更适合普通备份文件；若文件非常大，后续可改成 multipart
- `-verify=true` 只校验远端大小，不校验远端哈希
- 若 `-key` 为空，远端对象名会取本地文件名，如 `all-stock.db`