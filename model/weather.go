package model

import (
	"ImyouboHomeKit/common/enum"
	"ImyouboHomeKit/model/dao"
	xdb "ImyouboHomeKit/utils/db"
	"ImyouboHomeKit/utils/qweather"
)

func GetLatestWeather(locationId uint64, wt enum.WeatherInfoType) (*dao.LocationWeather, error){
	db, err := xdb.GetDefaultDB()
	if err != nil {
		return nil, err
	}
	return dao.GetOneLocationWeather(db, "location_id = ? and info_type = ?", locationId, byte(wt))
}

func InsertOrUpdateLocationWeather(locationId uint64, resp qweather.QWResp) (int64, error) {
	db, err := xdb.GetDefaultDB()
	if err != nil {
		return 0, err
	}
	updateTime, err := resp.GetUpdateTime()
	if err != nil {
		return 0, err
	}
	return dao.InsertOrUpdateLocationWeather(db, &dao.LocationWeather{
		LocationId: locationId,
		InfoType: byte(resp.GetType()),
		UpdateTime: updateTime,
		WeatherInfo: resp.String(),
	})
}
