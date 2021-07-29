package config

import "ImyouboHomeKit/utils"

const (
	SystemSection = "system"
	UrlSection = "url"
	DataSourceSection = "datasource"
)

type SystemConf struct {
	Port string `ini:"port"`
	ApiKey string `ini:"apikey"`
	UseDbCache bool `ini:"useDbCache"`
}

type UrlConf struct {
	GetRealtimeWeather string `ini:"getRealtimeWeather"`
	GetFutureDaysWeather string `ini:"getFutureDaysWeather"`
	GetFutureHourWeather string `ini:"getFutureHourWeather"`
	GetWeatherIndices string `ini:"getWeatherIndices"`
	GetRealtimeAir string `ini:"getRealtimeAir"`
}

type DataSourceConf struct {
	Url string `ini:"url"`
	Driver string `ini:"driver"`
}

var GlobalIniConfig *utils.IniConfig
var SystemConfig = new(SystemConf)
var UrlConfig = new(UrlConf)
var DataSourceConfig = new(DataSourceConf)


func InitConf(confPath string) error {
	if confPath == "" {
		confPath = "config.ini"
	}
	ini, err := utils.LoadIni(confPath)
	if err != nil {
		return err
	}
	GlobalIniConfig = ini
	return LoadConfigObj()
}

func LoadConfigObj() error {
	err := GlobalIniConfig.Conf.Section(SystemSection).MapTo(SystemConfig)
	if err != nil {
		return err
	}
	err = GlobalIniConfig.Conf.Section(UrlSection).MapTo(UrlConfig)
	if err != nil {
		return err
	}
	err = GlobalIniConfig.Conf.Section(DataSourceSection).MapTo(DataSourceConfig)
	return err
}
