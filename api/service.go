package api

import (
	"ImyouboHomeKit/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var HomeKitServiceSvc HomeKitService

type HomeKitService interface {
	ListLocationInfo(ctx *HttpContext, req *ListLocationInfoRequest) (*Response, error)
	GetRealtimeWeather(ctx *HttpContext) (*Response, error)
	GetFutureDaysWeather(ctx *HttpContext) (*Response, error)
	GetFutureHoursWeather(ctx *HttpContext) (*Response, error)
	GetWeatherIndices(ctx *HttpContext) (*Response, error)
	GetRealtimeAir(ctx *HttpContext) (*Response, error)
}

func homekitServiceListLocationInfo(c *gin.Context)  {
	p := new(ListLocationInfoRequest)
	if err := c.ShouldBind(&p); err != nil {
		handleResponse(c, nil, errors.RequestParamsFormatError(err))
		return
	}
	resp, err := HomeKitServiceSvc.ListLocationInfo(&HttpContext{c}, p)
	handleResponse(c, resp, err)
}

func homekitServiceGetRealtimeWeather(c *gin.Context) {
	resp, err := HomeKitServiceSvc.GetRealtimeWeather(&HttpContext{c})
	handleResponse(c, resp, err)
}

func homekitServiceGetWeatherIndices(c *gin.Context) {
	resp, err := HomeKitServiceSvc.GetWeatherIndices(&HttpContext{c})
	handleResponse(c, resp, err)
}

func homekitServiceGetRealtimeAir(c *gin.Context) {
	resp, err := HomeKitServiceSvc.GetRealtimeAir(&HttpContext{c})
	handleResponse(c, resp, err)
}

func homekitServiceGetFutureDaysWeather(c *gin.Context) {
	resp, err := HomeKitServiceSvc.GetFutureDaysWeather(&HttpContext{c})
	handleResponse(c, resp, err)
}

func homekitServiceGetFutureHoursWeather(c *gin.Context) {
	resp, err := HomeKitServiceSvc.GetFutureHoursWeather(&HttpContext{c})
	handleResponse(c, resp, err)
}


func handleResponse(c *gin.Context, resp *Response, err error) {
	if err != nil {
		c.JSON(http.StatusOK, ErrResp(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}
