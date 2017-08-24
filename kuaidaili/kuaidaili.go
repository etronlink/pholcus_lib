package pholcus_lib

// 基础包
import (
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	//. "github.com/henrylee2cn/pholcus/app/spider/common"    //选用
	"github.com/henrylee2cn/pholcus/common/goquery" //DOM解析
	//"github.com/henrylee2cn/pholcus/logs"                   //信息输出

	// net包
	//"net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	"strconv"
	//"strings"

	// 其他包
	"fmt"
	// "math"
	// "time"
	// "io/ioutil"
)

func init() {
	Kuaidaili.Register()
}

var Kuaidaili = &Spider{
	Name:         "获取免费代理列表",
	Description:  `获取免费代理列表 [www.kuaidaili.com]`,
	Pausetime:    2000,
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {

			var count = 10

			for i := 1; i <= count; i++ {
				ctx.AddQueue(&request.Request{
					//Url:          "http://www.kuaidaili.com/free/inha/" + strconv.Itoa(i) + "/",
					Url:          "http://www.kuaidaili.com/ops/proxylist/" + strconv.Itoa(i) + "/",
					Rule:         "代理列表",
					DownloaderID: 0,
				})
			}
		},

		Trunk: map[string]*Rule{
			"代理列表": {
				ItemFields: []string{
					"IP",
					"PORT",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					//fmt.Println(query.Find("tbody").Text())
					fmt.Println("url is " + ctx.Request.Url)
					query.Find("#freelist tbody tr").Each(func(i int, s *goquery.Selection) {
						ip := ""
						port := ""
						s.Find("td").Each(func(i int, si *goquery.Selection) {
							name, _ := si.Attr("data-title")
							//fmt.Println(name)
							if name == "IP" {
								ip = si.Text()
							}
							if name == "PORT" {
								port = si.Text()
							}
						})
						fmt.Println("ip: " + ip + " port: " + port)
						ctx.Output(map[int]interface{}{
							0: ip,
							1: port,
						})
					})
				},
			},
		},
	},
}

