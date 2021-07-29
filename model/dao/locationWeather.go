package dao

import (
	"ImyouboHomeKit/common/enum"
	"ImyouboHomeKit/utils"
	xdb "ImyouboHomeKit/utils/db"
	"fmt"
	"time"
)

const LocationWeatherTable = "location_weather"

type LocationWeather struct {
	LocationId  uint64    `json:"locationId" xorm:"location_id"`
	InfoType    byte      `json:"infoType" xorm:"info_type"`
	UpdateTime  time.Time `json:"updateTime" xorm:"update_time"`
	WeatherInfo string    `json:"weatherInfo" xorm:"weather_info"`
	Ct          time.Time `json:"ct" xorm:"ct"`
	Ut          time.Time `json:"ut" xorm:"ut"`
}

func GetOneLocationWeather(db *xdb.DB, where string, data ...interface{}) (*LocationWeather, error) {
	res := new(LocationWeather)
	if has, err := db.Table(LocationWeatherTable).Where(where, data...).Get(res); err != nil {
		return nil, err
	} else if !has {
		return nil, xdb.EmptyResultErr
	}
	defer db.Close()
	return res, nil
}

func InsertOrUpdateLocationWeather(db *xdb.DB, data *LocationWeather) (int64, error) {
	res, err := db.Exec(fmt.Sprintf("insert into "+LocationWeatherTable+"(location_id, info_type, update_time, weather_info) values(%d, %d, '%s', '%s') ON DUPLICATE KEY UPDATE info_type = values(info_type), update_time = values(update_time),  weather_info = values(weather_info)",
		data.LocationId, data.InfoType, data.UpdateTime.Format(utils.LayoutYYYYMMDDHHMMSS), data.WeatherInfo))
	if err != nil {
		return 0, err
	}
	defer db.Close()
	return res.RowsAffected()
}

func (v *LocationWeather) Expired() bool {
	switch enum.WeatherInfoType(v.InfoType) {
	case enum.WeatherInfo_WEATHER_REALTIME:
		fallthrough
	case enum.WeatherInfo_AIR_REALTIME:
		// 实时数据 10分钟失效
		return time.Now().Sub(v.UpdateTime) > 10 * time.Minute
	case enum.WeatherInfo_WEATHER_15DAYS:
		fallthrough
	case enum.WeatherInfo_INDICES_1DAYS:
		fallthrough
	case enum.WeatherInfo_WEATHER_24HOURS:
		//逐天或逐小时 1小时过期
		return time.Now().Sub(v.UpdateTime) > 1 * time.Hour
	}
	return false
}
