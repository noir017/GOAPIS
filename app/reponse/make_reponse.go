package reponse

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	C *gin.Context
}

type Basis struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BasisPage struct {
	Code      int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

// Response setting gin.JSON
func (g *Response) BaseResponse(data interface{}, intCode int, msg string) {
	g.C.JSON(200, Basis{
		Code: intCode,
		Msg:  msg,
		Data: data,
	})
	// switch errCode.(type) {
	// case int:
	// 	intCode := errCode.(int)

	// case string:
	// 	strCode := errCode.(string)
	// 	g.C.JSON(httpCode, Response{
	// 		Code: 9999,
	// 		Msg:  strCode,
	// 		Data: data,
	// 	})
	// }

	// return
}

// Response setting gin.JSON
func (g *Response) BasePageResponse(data interface{}, intCode int, msg string, total, totalPage int) {
	g.C.JSON(200, BasisPage{
		Code:      intCode,
		Msg:       msg,
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
}
