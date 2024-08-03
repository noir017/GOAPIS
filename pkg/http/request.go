package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	Text       string
	StatusCode int
}

type Requests struct {
	Client  resty.Client
	Headers map[string]string
	Timeout int
}

func NewRequests() *Requests {
	defaultHeaders := make(map[string]string)
	// defaultHeaders["Content-Type"] = "application/json"
	defaultHeaders["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0"
	req := Requests{
		Client:  *resty.New(),
		Headers: defaultHeaders,
		Timeout: 30,
	}
	transport, _ := req.Client.Transport()
	transport.TLSHandshakeTimeout = 30 * time.Second  // 增加tls握手时间
	req.Client.GetClient().Timeout = 30 * time.Second // 超时时间
	req.Client.SetHeaders(defaultHeaders)
	return &req
}

func (s Requests) Post(url string) *resty.Response {
	resp, err := s.Client.R().Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp
}
func (s Requests) Get(url string) (*Response, error) {

	resp, err := s.Client.R().EnableTrace().Get(url)
	var res Response
	if err != nil {
		fmt.Println("Error:", err)
		return &res, errors.New("请求错误")
	}
	res.StatusCode = resp.StatusCode()
	res.Text = resp.String()

	// respText := resp.String()
	// fmt.Println(respText)

	// traceInfo := resp.Request.TraceInfo()
	// fmt.Println("请求方法:", traceInfo)

	// var data map[string]interface{}
	// err2 := json.Unmarshal([]byte(respText), &data)
	// if err2 != nil || resp == nil {
	// 	fmt.Println("Error:", err2)
	// } else {
	// 	res.StatusCode = resp.StatusCode()
	// 	res.Text = resp.String()
	// }
	return &res, nil
}
