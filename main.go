package main

import (
	"ImyouboHomeKit/api"
	"ImyouboHomeKit/logic"
	"flag"
	"log"
)

var confPath string

func main() {
	// 处理命令行参数
	handlerArgs()
	// 创建http服务实例
	http := api.ProcHttp{}
	// 初始化并启动服务
	err := http.Init(confPath)
	if err != nil {
		log.Printf("Init http server error: %v", err)
		return
	}
	err = http.Start(new(logic.HomeKitServiceImpl))
	if err != nil {
		log.Printf("Http server start error: %v", err)
	}
}

func handlerArgs() {
	flag.StringVar(&confPath,"f", "config.ini", "指定配置文件路径")
	flag.String("h", "", "使用帮助")
	flag.Parse()
	if flag.Lookup("-h") != nil {
		flag.PrintDefaults()
	}
}