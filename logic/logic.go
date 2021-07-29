package logic

import (
	"ImyouboHomeKit/api"
	"ImyouboHomeKit/common/enum"
	"ImyouboHomeKit/config"
	"ImyouboHomeKit/errors"
	"ImyouboHomeKit/model"
	"ImyouboHomeKit/utils"
	"ImyouboHomeKit/utils/db"
	"ImyouboHomeKit/utils/qweather"
	"fmt"
	"strconv"
)

type HomeKitServiceImpl struct {}


func getQWeatherRespFromDbOrApi(refresh bool, wt enum.WeatherInfoType, locationId uint64, getQWRespFunc func() (qweather.QWResp, error), resp qweather.QWResp) (qweather.QWResp, error) {
	if refresh || config.SystemConfig.UseDbCache {
		return getNowWeatherAndUpdateDb(config.SystemConfig.UseDbCache, locationId, getQWRespFunc)
	} else {
		lastWeather, err := model.GetLatestWeather(locationId, wt)
		if err != nil && err != db.EmptyResultErr {
			return nil, errors.ErrDaoError(model.GetLatestWeather, err, "id -> %d, type -> %s", locationId, wt.String())
		} else if err == db.EmptyResultErr || lastWeather == nil || lastWeather.Expired() {
			return getNowWeatherAndUpdateDb(config.SystemConfig.UseDbCache, locationId, getQWRespFunc)
		} else {
			return resp, utils.JsonUnmarshal([]byte(lastWeather.WeatherInfo), resp)
		}
	}
}

func getNowWeatherAndUpdateDb(useDbCache bool, locationId uint64, getWeatherFunc func() (qweather.QWResp, error)) (qweather.QWResp, error) {
	resp, err := getWeatherFunc()
	if err != nil {
		return nil, err
	} else if err = qweather.CheckRespCodeSuccess(resp);err != nil {
		return nil, err
	}
	if useDbCache {
		_, err = model.InsertOrUpdateLocationWeather(locationId, resp)
		if err != nil {
			return nil, errors.ErrDaoError(model.InsertOrUpdateLocationWeather, err,"id -> %d", locationId)
		}
	}
	return resp, err
}

func getLocationFromHttpContext(ctx *api.HttpContext) (location string, locationId uint64, err error) {
	location = ctx.Query("location")
	if location == "" {
		return "", 0, errors.InvalidRequestParam("location", location)
	}
	if locationId, err = strconv.ParseUint(location, 10, 64); err != nil {
		return location, 0, fmt.Errorf("parse int error: %v, location: %s", err, location)
	}
	_, err = model.GetLocationById(locationId)
	if err != nil {
		if err == db.EmptyResultErr {
			return location, locationId, fmt.Errorf("location id: %d not exists", locationId)
		}
		return location, locationId, errors.ErrDaoError(model.GetLocationById, err, "id -> %d", locationId)
	}
	return
}