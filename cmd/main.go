package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhuyanxi/CarnoFinance/pkg/api"
	"github.com/zhuyanxi/CarnoFinance/pkg/db"
	"github.com/zhuyanxi/CarnoFinance/pkg/domain"
	"github.com/zhuyanxi/CarnoFinance/pkg/helper"
	"github.com/zhuyanxi/CarnoFinance/pkg/xueqiu"
)

// curl --location 'http://localhost:8080/etf/rsrslist?period=25&date=20240705'

// flow:
// init history data: get all stock index and data

func main() {
	a := time.Unix(1709481600, 0).UTC().Format(helper.DateFormat)
	logrus.Infoln(a)
	b := time.Unix(1709481600, 0).Local().Format(helper.DateFormat)
	logrus.Infoln(b)

	envs := os.Environ()
	logrus.Infof("%+v", envs)

	ctx := context.Background()
	sqlite := db.NewSqlite(ctx)
	xqc := xueqiu.New()
	app := domain.NewDomain(ctx, sqlite, xqc)
	app.Init()

	r := gin.Default()
	// CORS for http://localhost:5173 origins
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:8080"
		// },
		MaxAge: 12 * time.Hour,
	}))
	frontendDistPath := filepath.Join("..", "ui", "build")
	r.Static("/_ui", frontendDistPath)
	r.NoRoute(func(c *gin.Context) {
		// 检查请求路径是否像是静态文件 (例如，以 .js, .css, .png 等结尾)
		// 避免将所有请求都重定向到 index.html，导致静态资源无法加载
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/assets/") || strings.HasPrefix(path, "/_app/") ||
			strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".css") ||
			strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpg") ||
			strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".gif") ||
			strings.HasSuffix(path, ".svg") || strings.HasSuffix(path, ".ico") {
			c.File(filepath.Join(frontendDistPath, path)) // 尝试直接提供静态文件
			return
		}

		// 对于所有其他未匹配的路由（即前端路由），返回 index.html
		c.File(filepath.Join(frontendDistPath, "index.html"))
	})

	r.GET("/healthz", func(ctx *gin.Context) {
		helper.GinOK(ctx)
	})

	//SZ159915
	r.POST("/etf/last", api.SetETFPrice(app))
	r.GET("/etf/rsrslist", api.GetRSRSList(app))
	r.GET("/etf/:code/close", api.GetETFClosePrices(app))

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// 假设你有一个包含日期、净资产、出金和入金的数据文件（例如CSV文件），你可以使用Go语言读取数据并计算真实年化收益率。
// 首先，我们假设你的数据文件名为 `data.csv`，其格式如下：

// ```
// Date,NetAssetValue,Deposit,Withdrawal
// 2013-01-01,10000,100,50
// 2013-01-02,10100,200,100
// ...
// ```

// 接下来是Go语言代码：

// ```go
// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	file, err := os.Open("data.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if len(records) < 2 {
// 		log.Fatal("Not enough data in CSV file")
// 	}

// 	var (
// 		previousNAV float64
// 		totalReturn float64 = 1.0
// 		startDate   time.Time
// 		endDate     time.Time
// 	)

// 	for i, record := range records[1:] {
// 		date, err := time.Parse("2006-01-02", record[0])
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		NAV, err := strconv.ParseFloat(record[1], 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		deposit, err := strconv.ParseFloat(record[2], 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		withdrawal, err := strconv.ParseFloat(record[3], 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if i == 0 {
// 			previousNAV = NAV
// 			startDate = date
// 			continue
// 		}

// 		adjustedNAV := previousNAV + deposit - withdrawal
// 		dailyReturn := (NAV - adjustedNAV) / adjustedNAV
// 		totalReturn *= (1 + dailyReturn)

// 		previousNAV = NAV
// 		endDate = date
// 	}

// 	cumulativeReturn := totalReturn - 1.0
// 	days := endDate.Sub(startDate).Hours() / 24
// 	annualizedReturn := math.Pow(totalReturn, 365/days) - 1.0

// 	fmt.Printf("Cumulative Return: %.2f%%\n", cumulativeReturn*100)
// 	fmt.Printf("Annualized Return: %.2f%%\n", annualizedReturn*100)
// }
// ```

// 这个Go程序执行以下步骤：

// 1. **读取CSV文件**：打开并读取CSV文件，假定第一行是标题行。
// 2. **解析数据**：逐行解析日期、净资产、入金和出金数据。
// 3. **计算每日收益率**：使用前一天的净资产和当天的入金、出金数据计算每日收益率。
// 4. **计算累积收益率**：累积每日收益率。
// 5. **计算年化收益率**：根据累积收益率和时间跨度计算年化收益率。

// ### 解释

// - **文件操作**：`os.Open` 打开CSV文件，`csv.NewReader` 和 `ReadAll` 读取文件内容。
// - **时间解析**：`time.Parse` 将字符串日期解析为 `time.Time` 类型。
// - **字符串转换**：`strconv.ParseFloat` 将字符串解析为浮点数。
// - **累积收益率计算**：逐步累积每日收益率。
// - **时间跨度计算**：通过计算起始日期和结束日期的差异来确定总天数。
// - **年化收益率计算**：使用数学公式计算年化收益率。

// 这个Go程序假定数据文件名为 `data.csv` 并且包含合适的数据格式。你可以根据实际情况修改文件名和路径。
