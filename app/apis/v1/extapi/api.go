package extapi

// 获取每日新闻接口
import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/reponse"
	"github.com/noir017/goapis/pkg/http"
	"github.com/noir017/goapis/pkg/tools"
	"github.com/sirupsen/logrus"
)

type ExtApi struct{}

func (s ExtApi) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello, Gin",
	})
}

// 获取每日新闻api
func (s ExtApi) GetNewsEveryDay(c *gin.Context) {
	response := reponse.Response{C: c}
	req := http.NewRequests()
	res, resErr := req.Get("https://api.vvhan.com/api/hotlist/wbHot")
	if resErr != nil {
		response.BaseResponse(nil, reponse.ERROR, "fetch failed")
		return
	}
	resJson, jsonErr := tools.StrToJson(res.Text)
	if jsonErr != nil {
		fmt.Println(jsonErr)

		response.BaseResponse(nil, reponse.ERROR, "jsonify failed")
		return
	}
	resData, dataErr := ParseVvhan(resJson)
	if dataErr != nil {
		response.BaseResponse(nil, reponse.ERROR, dataErr.Error())
		return
	}

	// SaveHotwords(resData)
	fmt.Println(resData, dataErr)
	response.BaseResponse(resData, reponse.SUCCESS, reponse.GetMsg(reponse.SUCCESS))

}
func fetchNewsEveryDay(url string) []Hotword {
	req := http.NewRequests()
	res, resErr := req.Get(url)

	logrus.Info(fmt.Sprintf("访问接口：%s", url))
	time.Sleep(time.Second * 2)
	if resErr != nil {
		logrus.Warn(fmt.Sprintf("接口【%s】访问错误："+resErr.Error(), url))
		return []Hotword{}
	}
	resJson, jsonErr := tools.StrToJson(res.Text)
	if jsonErr != nil {
		logrus.Warn(fmt.Sprintf("接口【%s】jsonify失败："+jsonErr.Error(), url))
		return []Hotword{}
	}
	resData, dataErr := ParseVvhan(resJson)
	if dataErr != nil {
		logrus.Warn(fmt.Sprintf("接口【%s】数据解析失败："+dataErr.Error(), url))
		return []Hotword{}
	}
	return resData
}

// 获取每日新闻
func SaveNewsEveryDay() {
	urls := []string{
		// "https://api.vvhan.com/api/hotlist/wbHot",
		// "https://api.vvhan.com/api/hotlist/toutiao",
		// "https://api.vvhan.com/api/hotlist/huPu",
		// "https://api.vvhan.com/api/hotlist/zhihuHot",
		// "https://api.vvhan.com/api/hotlist/bili",
		// "https://api.vvhan.com/api/hotlist/36Ke",
		// "https://api.vvhan.com/api/hotlist/baiduRD",
		// "https://api.vvhan.com/api/hotlist/douyinHot",
		"https://api.vvhan.com/api/hotlist/douban",
		"https://api.vvhan.com/api/hotlist/itNews",
		// "https://api.vvhan.com/api/hotlist/huXiu",
		// "https://api.vvhan.com/api/hotlist/woShiPm",
	}
	data := []Hotword{}
	for _, url := range urls {
		temp := fetchNewsEveryDay(url)
		data = append(data, temp...)
	}
	SaveHotwords(data)
	logrus.Info("每日新闻访问完毕，已储存")

}

// 获取数据库内随机新闻
func (s ExtApi) GetRandomNews(c *gin.Context) {
	funName := "SaveNewsEveryDay"
	isExecuted := CheckExecuted(funName)
	if !isExecuted {
		RecordExecuted(funName)
		go SaveNewsEveryDay()
	}
	response := reponse.Response{C: c}
	resData := GetRandomHotwords()
	response.BaseResponse(resData, reponse.SUCCESS, reponse.GetMsg(reponse.SUCCESS))
	// data, err := tools.ParseTenapicn(resJson)
	// fmt.Println(data, err)
	// req.Get("https://api-hot.efefee.cn/bilibili")

	// c.JSON(200, gin.H{
	// 	"message": "Hello, GetNewsEveryDay",
	// })
}
func (s ExtApi) DuolicateNews(c *gin.Context) {
	Duolicate()
	response := reponse.Response{C: c}
	response.BaseResponse(nil, reponse.SUCCESS, "删除成功")
}
