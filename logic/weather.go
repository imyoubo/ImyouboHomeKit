package logic

import (
	"ImyouboHomeKit/api"
	"ImyouboHomeKit/common/enum"
	"ImyouboHomeKit/errors"
	"ImyouboHomeKit/model"
	"ImyouboHomeKit/utils/qweather"
	"fmt"
	"strconv"
)

func (hsi *HomeKitServiceImpl) GetRealtimeWeather(ctx *api.HttpContext) (*api.Response, error) {
	location, locationId, err := getLocationFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}
	var weatherResp qweather.QWResp
	weatherResp = new(qweather.RealtimeWeatherResp)
	weatherResp, err = getQWeatherRespFromDbOrApi(ctx.Query("refresh") == "true",
		enum.WeatherInfo_WEATHER_REALTIME,
		locationId, qweather.GetRealtimeWeatherFunc(location), weatherResp)
	if err != nil {
		return nil, err
	}
	return api.SuccessDataResp(weatherResp.FormatTime())
}


func (hsi *HomeKitServiceImpl) ListLocationInfo(ctx *api.HttpContext, req *api.ListLocationInfoRequest) (*api.Response, error) {
	list, total, err := model.ListLocationInfo(req)
	if err != nil {
		return nil, errors.ErrDaoError(model.ListLocationInfo, err, "param -> %v", req)
	}
	return api.SuccessDataResp(api.BuildListItems(total, req.Offset, req.Limit, len(list), list))
}

func (hsi *HomeKitServiceImpl) GetFutureDaysWeather(ctx *api.HttpContext) (*api.Response, error) {
	location, locationId, err := getLocationFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}
	d := ctx.DefaultQuery("d", "3")
	days, err := strconv.ParseUint(d, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse int error: %v, d: %s", err, d)
	} else if days > qweather.DefaultWeatherDays {
		return nil, fmt.Errorf("days cannot exceed %d", qweather.DefaultWeatherDays)
	}
	var weatherResp qweather.QWResp
	weatherResp = new(qweather.FutureDaysResp)
	weatherResp, err = getQWeatherRespFromDbOrApi(ctx.Query("refresh") == "true",
		enum.WeatherInfo_WEATHER_15DAYS,
		locationId, qweather.GetDailyWeatherFunc(location, qweather.DefaultWeatherDays), weatherResp)
	if err != nil {
		return nil, err
	}
	return api.SuccessDataResp(weatherResp.Limit(uint(days)).FormatTime())
}

func (hsi *HomeKitServiceImpl) GetFutureHoursWeather(ctx *api.HttpContext) (*api.Response, error) {
	location, locationId, err := getLocationFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}
	h := ctx.DefaultQuery("h", "12")
	hour, err := strconv.ParseUint(h, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse int error: %v, h: %s", err, h)
	} else if hour > qweather.DefaultWeatherHours {
		return nil, fmt.Errorf("hour cannot exceed %d", qweather.DefaultWeatherHours)
	}
	var weatherResp qweather.QWResp
	weatherResp = new(qweather.FutureHoursWeatherResp)
	weatherResp, err = getQWeatherRespFromDbOrApi(ctx.Query("refresh") == "true",
		enum.WeatherInfo_WEATHER_24HOURS,
		locationId, qweather.GetHourlyWeatherFunc(location, qweather.DefaultWeatherHours), weatherResp)
	if err != nil {
		return nil, err
	}
	return api.SuccessDataResp(weatherResp.Limit(uint(hour)).FormatTime())
}

func (hsi *HomeKitServiceImpl) GetWeatherIndices(ctx *api.HttpContext) (*api.Response, error) {
	location, locationId, err := getLocationFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}
	d := ctx.DefaultQuery("d", "1")
	days, err := strconv.ParseUint(d, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse int error: %v, h: %s", err, d)
	} else if days > qweather.DefaultIndicesDays {
		return nil, fmt.Errorf("days cannot exceed %d", qweather.DefaultIndicesDays)
	}
	itp := ctx.Query("type")
	var weatherResp qweather.QWResp
	weatherResp = new(qweather.WeatherIndicesResp)
	weatherResp, err = getQWeatherRespFromDbOrApi(ctx.Query("refresh") == "true",
		enum.WeatherInfo_INDICES_1DAYS,
		locationId, qweather.GetWeatherIndicesFunc(location, qweather.DefaultIndicesDays, "0"), weatherResp)
	if err != nil {
		return nil, err
	}
	return api.SuccessDataResp((weatherResp.Limit(uint(days))).(*qweather.WeatherIndicesResp).FilterType(itp).FormatTime())
}

func (hsi *HomeKitServiceImpl) GetRealtimeAir(ctx *api.HttpContext) (*api.Response, error) {
	location, locationId, err := getLocationFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}
	var weatherResp qweather.QWResp
	weatherResp = new(qweather.RealtimeAirResp)
	weatherResp, err = getQWeatherRespFromDbOrApi(ctx.Query("refresh") == "true",
		enum.WeatherInfo_AIR_REALTIME,
		locationId, qweather.GetRealtimeAirFunc(location), weatherResp)
	if err != nil {
		return nil, err
	}
	return api.SuccessDataResp(weatherResp.FormatTime())
}