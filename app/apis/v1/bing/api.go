package bing

// 获取每日新闻接口
import (
	"github.com/gin-gonic/gin"
	"github.com/noir017/goapis/app/reponse"
)

type BingApi struct{}

func (s BingApi) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello, Gin",
	})
}

// 获取每日新闻api
func (s BingApi) RecordReward(c *gin.Context) {
	bingMes := c.Query("mes")
	response := reponse.Response{C: c}
	if bingMes == "" {
		response.BaseResponse("无请求参数", reponse.SUCCESS, reponse.GetMsg(reponse.SUCCESS))
	} else {
		temp := Rewards{Detail: bingMes}
		SaveRewards(temp)
		response.BaseResponse("保存成功", reponse.SUCCESS, reponse.GetMsg(reponse.SUCCESS))
	}

}
