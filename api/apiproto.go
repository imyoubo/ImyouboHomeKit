package api

import "github.com/gin-gonic/gin"

const (
	Success = "success"
)

// HttpServer is the http server, Create an instance of GinServer, by using NewGinServer()
type HttpServer struct {
	*gin.Engine
}

// HttpContext warp gin HttpContext
type HttpContext struct {
	*gin.Context
}

type ResponseData struct {
	Ent interface{} `json:"ent,omitempty"`
	Ext interface{} `json:"ext,omitempty"`
}

type Response struct {
	Ret  int           `json:"ret"`
	Msg  string        `json:"msg,omitempty"`
	Data *ResponseData `json:"data,omitempty"`
}

type ListItems struct {
	Items  interface{} `json:"items"`
	Total  int64       `json:"total"`
	Offset int         `json:"offset"`
	More   bool        `json:"more"`
}

func BuildListItems(total int64, offset, limit, len int, items interface{}) *ListItems {
	return &ListItems{
		Items:  items,
		Total:  total,
		Offset: offset + len,
		More:   total > int64(offset+limit),
	}
}

func SuccessResp() (*Response, error) {
	return FmtResp(1, Success, nil), nil
}

func SuccessDataResp(data interface{}) (*Response, error) {
	return FmtResp(1, Success, data), nil
}

func SuccessIfNotError(err error) (*Response, error) {
	if err != nil {
		return nil, err
	}
	return SuccessResp()
}

func ErrResp(err error) *Response {
	return FmtResp(-1, err.Error(), nil)
}

func FmtResp(code int, msg string, data interface{}) *Response {
	if code == 0 {
		code = 1
	}
	if msg == "" {
		msg = Success
	}
	if data == nil {
		return &Response{
			Ret: code,
			Msg: msg,
		}
	} else {
		return &Response{
			Ret: code,
			Msg: msg,
			Data: &ResponseData{
				Ent: data,
			},
		}
	}
}
