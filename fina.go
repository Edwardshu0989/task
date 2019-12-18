package main

import (
	// "crypto/tls"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"task/lib/logger"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"github.com/yeeyuntech/yeego"
	"gitlab.yeeyuntech.com/yee/easyweb"
	"gitlab.yeeyuntech.com/yee/easyweb_cms"
)

const (
	google             string = "play.google.com"
	comName            string = "Apps on Google Play"
	apk                string = "apk"
	CommodityStateDown        = "down" // 已下架
	limit              int64  = 1      // 查询条数
)

type Commodity struct {
	Id                   int64   `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 主键
	Name                 string  `json:"name"`                                 // 商品名称(30)
	DownloadUrl          string  `json:"download_url"`                         // 下载链接地址
	Slogans              string  `json:"slogans"`                              // 广告语(30)
	Logo                 string  `json:"logo"`                                 // logo图片
	Tag                  string  `json:"tag"`                                  // 广告标签图片
	Description          string  `gorm:"type:text" json:"description"`         // 简介(300)
	Procedure            string  `gorm:"type:text" json:"procedure"`           // 贷款流程(1000)
	Sort                 int64   `json:"sort"`                                 // 排序ASC
	StartTime            int64   `json:"start_time"`                           // 开始时间
	EndTime              int64   `json:"end_time"`                             // 结束时间
	LendingTime          int64   `json:"lending_time"`                         // 放款时间(分钟为单位)
	LoanableAmountMin    int64   `json:"loanable_amount_min"`                  // 最小可贷款金额
	LoanableAmountMax    int64   `json:"loanable_amount_max"`                  // 最大可贷款金额
	MonthInterestRateMin float64 `json:"month_interest_rate_min"`              // 最小月利率
	MonthInterestRateMax float64 `json:"month_interest_rate_max"`              // 最大月利率
	MonthServiceFee      float64 `json:"month_service_fee"`                    // 月服务费
	LendingCycle         string  `json:"lending_cycle"`                        // 借贷周期(eg:1,2,24 表示支持1月，2月，24月)
	State                string  `json:"state"`                                // 状态
	Remark               string  `json:"remark"`                               //备注
	CreateAt             int64   `json:"create_at"`
	UpdateAt             int64   `json:"update_at"`
}

func (Commodity) TableName() string {
	return "lm_commodity"
}

var (
	defaultDb     *gorm.DB // 默认数据库
	code          int
	count         int
	redirect      string
	err           error
	title         string
	commodityName string
	state         string
	c             = colly.NewCollector()
	ct            Commodity
	url           string
)

func init() {
	yeego.MustInitConfig(".", "conf")
	var (
		logCfg = logger.LogConfig{
			DebugPath: "logs/debug.log",
			ErrorPath: "logs/error.log",
		}
	)
	logger.InitLogger(logCfg)
	dbConf := easyweb_cms.DbConfig{
		UserName: yeego.Config.GetString("database.UserName"),
		Password: yeego.Config.GetString("database.Password"),
		Host:     yeego.Config.GetString("database.Host"),
		Port:     yeego.Config.GetString("database.Port"),
		DbName:   yeego.Config.GetString("database.DbName"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		easyweb.Logger.Error("数据库连接失败:%s", err.Error())
	}
	db.DB().SetMaxIdleConns(2000)
	defaultDb = db
}

func main() {
	d := cron.New()
	d.AddFunc("0 0/30 * * * *", task)
	d.Start()
	t1 := time.NewTimer(time.Minute * 1)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Minute * 1)
		}
	}
}

func task() {
	// defaultDb.Model(Commodity{}).Count(&count)
	if cts, err := FetchApi(0, 1); err != nil {
		easyweb.Logger.Error("获取线上地址出错:%s", err.Error())
	} else {
		var dast map[string]interface{}
		if err := json.Unmarshal(cts, &dast); err != nil {
			easyweb.Logger.Error("解析线上地址信息出错:%s", err.Error())
			return
		}
		mapinfos := dast["data"].(map[string]interface{})
		fmt.Println(count)
		count, _ = mapinfos["count"].(int)
	}
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Minute)
		page := i
		pagesize := 1
		if ct, err := FetchApi(page, pagesize); err != nil {
			easyweb.Logger.Error("获取线上地址出错:%s", err.Error())
		} else {
			var dat map[string]interface{}
			if err := json.Unmarshal(ct, &dat); err != nil {
				easyweb.Logger.Error("解析线上地址信息出错:%s", err.Error())
				return
			}
			mapinfo := dat["data"].(map[string]interface{})
			commodity := (mapinfo["data"].([]interface{}))[0].(map[string]interface{})
			// defaultDb.Limit(limit).Offset(i).First(&ct)
			// url = ct.DownloadUrl
			// commodityName = ct.Name
			url = commodity["download_url"].(string)
			commodityName = commodity["name"].(string)
			if strings.Index(url, "https") == -1 {
				url = strings.Replace(url, "http", "https", 1)
			}
			fmt.Println(url)
			if strings.Contains(url, google) {
				download(url, commodityName)
			} else {
				state = "on"
				for state == "on" {
					redirect, err := download2(url)
					if err != nil {
						break
					}
					fmt.Println(redirect)
					switch {
					case strings.Contains(redirect, google):
						download(url, commodityName)
						state = "off"
					case strings.Contains(redirect, apk):
						easyweb.Logger.Error("该链接不是谷歌市场链接%s:", url)
						state = "off"
					default:
						state = "on"
					}
				}
			}
		}
	}
}

func FetchApi(page, pagesize int) ([]byte, error) {
	pages := strconv.Itoa(page)
	pagesizes := strconv.Itoa(pagesize)
	var URL = "http://data.loanmarket123.com/v1/commodity/list?page=" + pages + "&pagesize=" + pagesizes
	var commodities []byte
	if response, err := http.Get(URL); err != nil {
		easyweb.Logger.Error("获取线上地址出错:%s", err.Error())
		return nil, err
	} else {
		if commodities, err = ioutil.ReadAll(response.Body); err != nil {
			easyweb.Logger.Error("读取线上地址信息出错:%s", err.Error())
			return nil, err
		}
	}
	return commodities, err
}

func download(url string, commodity string) {
	title, err := BrowserAccess(url)
	if err != nil {
		easyweb.Logger.Error("链接%s谷歌链接访问出错:%s", url, err.Error())
		defaultDb.Model(Commodity{}).Where("name = ?", commodity).Update(Commodity{State: CommodityStateDown, Remark: "未在谷歌市场检测到APP"})
	} else {
		easyweb.Logger.Error("111111链接%s谷歌链接访问成功:%s", url, title)
		// if !strings.Contains(title, comName) {
		// 	defaultDb.Model(Commodity{}).Where("name = ?", commodity).Update(Commodity{State: CommodityStateDown, Remark: "未在谷歌市场检测到APP"})
		// }
	}
}

func BrowserAccess(url string) (string, error) {
	var buf []byte
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 27*time.Second)
	defer cancel()
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#ZCHFDb`),
		chromedp.Text(`#fcxH9b > div.WpDbMd > c-wiz > div > div.ZfcPIb > div > div.JNury.Ekdcne > div > c-wiz:nth-child(1) > c-wiz:nth-child(1) > div > div.D0ZKYe > div > div.sIskre > c-wiz:nth-child(1) > h1 > span`, &example),
		chromedp.CaptureScreenshot(&buf),
	)
	return example, err
}

func download2(url string) (string, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		easyweb.Logger.Error("http get error:%s", err.Error())
		return "1", err
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()
	if len(resp.Header["Location"]) > 0 {
		return resp.Header["Location"][0], nil
	} else {
		return "1", fmt.Errorf("该链接不是谷歌链接")
	}

}
