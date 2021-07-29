package api

import (
	"ImyouboHomeKit/config"
	"ImyouboHomeKit/errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ProcHttp struct{}

type Processor interface {
	Init(string) error
	Start(HomeKitService) error
}

func (p *ProcHttp) Init(path string) error {
	log.Printf("Init gin http server, conf path: %s -->", path)
	err := config.InitConf(path)
	if err != nil {
		return errors.InitConfigError(err)
	}
	return nil
}

func (p *ProcHttp) Start(service HomeKitService) error {
	serv := NewHttpServer()
	RegisterHttpServer(serv, service)
	return serv.Run("0.0.0.0:" + config.SystemConfig.Port)
}

func NewHttpServer() *HttpServer {
	// 实例化gin Server
	router := gin.Default()
	return &HttpServer{router}
}

func RegisterHttpServer(e *HttpServer, service HomeKitService)  {
	HomeKitServiceSvc = service
	// 服务探活
	e.GET("/ping", func(c *gin.Context) {c.JSON(http.StatusOK, FmtResp(1, Success, "pong"))})
	// 分页查询地区信息
	e.POST("/api/weather/location/list", homekitServiceListLocationInfo)
	// 实时天气接口
	e.GET("/api/weather/now", homekitServiceGetRealtimeWeather)
	// 获取未来逐小时天气
	e.GET("/api/weather/hours", homekitServiceGetFutureHoursWeather)
	// 获取未来逐天天气
	e.GET("/api/weather/days", homekitServiceGetFutureDaysWeather)
	// 获取天气生活指数
	e.GET("/api/weather/indices", homekitServiceGetWeatherIndices)
	// 获取实时空气质量
	e.GET("/api/weather/air/now", homekitServiceGetRealtimeAir)

}