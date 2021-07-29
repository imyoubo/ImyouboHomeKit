package qweather

import (
	"ImyouboHomeKit/config"
	"ImyouboHomeKit/utils"
	"fmt"
)

const (
	DefaultWeatherDays = 15
	DefaultWeatherHours = 24
	DefaultIndicesDays = 1
)

func CheckRespCodeSuccess(resp QWResp) error {
	code := resp.GetCode()
	if code != "200" {
		return fmt.Errorf("%s return error code: %s", resp.GetType().String(), code)
	}
	return nil
}

func GetRealtimeWeather(location string) (*RealtimeWeatherResp, error) {
	res := new(RealtimeWeatherResp)
	url := fmt.Sprintf(config.UrlConfig.GetRealtimeWeather+"?key=%s&location=%s", config.SystemConfig.ApiKey, location)
	resp, err := utils.NewHttpClient(nil).Get(url)
	if err != nil {
		return nil, err
	}
	return res, utils.JsonUnmarshal(resp, res)
}

func GetDailyWeather(location string, days uint64) (*FutureDaysResp, error) {
	res := new(FutureDaysResp)
	url := fmt.Sprintf(config.UrlConfig.GetFutureDaysWeather+"?key=%s&location=%s", days, config.SystemConfig.ApiKey, location)
	resp, err := utils.NewHttpClient(nil).Get(url)
	if err != nil {
		return nil, err
	}
	return res, utils.JsonUnmarshal(resp, res)
}

func GetHourlyWeather(location string, hours uint64) (*FutureHoursWeatherResp, error) {
	res := new(FutureHoursWeatherResp)
	url := fmt.Sprintf(config.UrlConfig.GetFutureHourWeather+"?key=%s&location=%s", hours, config.SystemConfig.ApiKey, location)
	resp, err := utils.NewHttpClient(nil).Get(url)
	if err != nil {
		return nil, err
	}
	return res, utils.JsonUnmarshal(resp, res)
}

func GetWeatherIndices(location string, days uint64, tp string) (*WeatherIndicesResp, error) {
	res := new(WeatherIndicesResp)
	url := fmt.Sprintf(config.UrlConfig.GetWeatherIndices+"?key=%s&location=%s&type=%s", days, config.SystemConfig.ApiKey, location, tp)
	resp, err := utils.NewHttpClient(nil).Get(url)
	if err != nil {
		return nil, err
	}
	return res, utils.JsonUnmarshal(resp, res)
}

func GetRealtimeAir(location string) (*RealtimeAirResp, error) {
	res := new(RealtimeAirResp)
	url := fmt.Sprintf(config.UrlConfig.GetRealtimeAir+"?key=%s&location=%s", config.SystemConfig.ApiKey, location)
	resp, err := utils.NewHttpClient(nil).Get(url)
	if err != nil {
		return nil, err
	}
	return res, utils.JsonUnmarshal(resp, res)
}

func GetRealtimeWeatherFunc(location string) func() (QWResp, error) {
	return func() (QWResp, error) {
		return GetRealtimeWeather(location)
	}
}

func GetDailyWeatherFunc(location string, days uint64) func() (QWResp, error)  {
	return func() (QWResp, error) {
		return GetDailyWeather(location, days)
	}
}

func GetRealtimeAirFunc(location string) func() (QWResp, error)  {
	return func() (QWResp, error) {
		return GetRealtimeAir(location)
	}
}

func GetHourlyWeatherFunc(location string, hours uint64) func() (QWResp, error)  {
	return func() (QWResp, error) {
		return GetHourlyWeather(location, hours)
	}
}

func GetWeatherIndicesFunc(location string, days uint64, tp string)  func() (QWResp, error)  {
	return func() (QWResp, error) {
		return GetWeatherIndices(location, days, tp)
	}
}