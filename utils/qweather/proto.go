package qweather

import (
	"ImyouboHomeKit/common/enum"
	"ImyouboHomeKit/utils"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type QWResp interface {
	GetCode() string
	GetType() enum.WeatherInfoType
	String() string
	GetUpdateTime() (time.Time, error)
	FormatTime() QWResp
	Limit(uint) QWResp
}

type RealtimeWeatherRespNow struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type RealtimeWeatherResp struct {
	Code       string                  `json:"code"`
	UpdateTime string                  `json:"updateTime"`
	FxLink     string                  `json:"fxLink"`
	Now        *RealtimeWeatherRespNow `json:"now"`
}

func (v *RealtimeWeatherResp) GetUpdateTime() (time.Time, error) {
	return utils.ParseTimeFromIOS8601(v.UpdateTime)
}

func (v *RealtimeWeatherResp) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(bytes)
}

func (v *RealtimeWeatherResp) FormatTime() QWResp {
	v.UpdateTime = utils.ConvertTimeFromIOS8601(v.UpdateTime)
	v.Now.ObsTime = utils.ConvertTimeFromIOS8601(v.Now.ObsTime)
	return v
}

func (v *RealtimeWeatherResp) GetCode() string {
	return v.Code
}

func (v *RealtimeWeatherResp) GetType() enum.WeatherInfoType {
	return enum.WeatherInfo_WEATHER_REALTIME
}

func (v *RealtimeWeatherResp) Limit(limit uint) QWResp {
	return v
}

type FutureDaysRespDaily struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	MoonSet        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	UvIndex        string `json:"uvIndex"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	Dew            string `json:"dew"`
}

type FutureDaysResp struct {
	Code       string                 `json:"code"`
	UpdateTime string                 `json:"updateTime"`
	FxLink     string                 `json:"fxLink"`
	Daily      []*FutureDaysRespDaily `json:"daily"`
}

func (v *FutureDaysResp) GetCode() string {
	return v.Code
}

func (v *FutureDaysResp) GetType() enum.WeatherInfoType {
	return enum.WeatherInfo_WEATHER_15DAYS
}

func (v *FutureDaysResp) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(bytes)
}

func (v *FutureDaysResp) GetUpdateTime() (time.Time, error) {
	return utils.ParseTimeFromIOS8601(v.UpdateTime)
}

func (v *FutureDaysResp) FormatTime() QWResp {
	v.UpdateTime = utils.ConvertTimeFromIOS8601(v.UpdateTime)
	return v
}

func (v *FutureDaysResp) Limit(limit uint) QWResp {
	v.Daily = v.Daily[:limit]
	return v
}

type FutureHoursWeatherRespHourly struct {
	FxTime    string `json:"fxTime"`
	Temp      string `json:"temp"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pop       string `json:"pop"`
	Pressure  string `json:"pressure"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type FutureHoursWeatherResp struct {
	Code       string                          `json:"code"`
	UpdateTime string                          `json:"updateTime"`
	FxLink     string                          `json:"fxLink"`
	Hourly     []*FutureHoursWeatherRespHourly `json:"hourly"`
}

func (v *FutureHoursWeatherResp) GetCode() string {
	return v.Code
}

func (v *FutureHoursWeatherResp) GetType() enum.WeatherInfoType {
	return enum.WeatherInfo_WEATHER_24HOURS
}

func (v *FutureHoursWeatherResp) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(bytes)
}

func (v *FutureHoursWeatherResp) GetUpdateTime() (time.Time, error) {
	return utils.ParseTimeFromIOS8601(v.UpdateTime)
}

func (v *FutureHoursWeatherResp) FormatTime() QWResp {
	v.UpdateTime = utils.ConvertTimeFromIOS8601(v.UpdateTime)
	for _, h := range v.Hourly {
		h.FxTime = utils.ConvertTimeFromIOS8601(h.FxTime)
	}
	return v
}

func (v *FutureHoursWeatherResp) Limit(limit uint) QWResp {
	v.Hourly = v.Hourly[:limit]
	return v
}

type WeatherIndicesRespDaily struct {
	Date     string `json:"date"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Text     string `json:"text"`
}

type WeatherIndicesResp struct {
	Code       string                     `json:"code"`
	UpdateTime string                     `json:"updateTime"`
	FxLink     string                     `json:"fxLink"`
	Daily      []*WeatherIndicesRespDaily `json:"daily"`
}

func (v *WeatherIndicesResp) GetCode() string {
	return v.Code
}

func (v *WeatherIndicesResp) GetType() enum.WeatherInfoType {
	return enum.WeatherInfo_INDICES_1DAYS
}

func (v *WeatherIndicesResp) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(bytes)
}

func (v *WeatherIndicesResp) GetUpdateTime() (time.Time, error) {
	return utils.ParseTimeFromIOS8601(v.UpdateTime)
}

func (v *WeatherIndicesResp) FormatTime() QWResp {
	v.UpdateTime = utils.ConvertTimeFromIOS8601(v.UpdateTime)
	return v
}

func (v *WeatherIndicesResp) Limit(limit uint) QWResp {
	var daily = make([]*WeatherIndicesRespDaily, 0)
	day := 24 * time.Hour
	nowDate, _ := time.Parse(utils.LayoutYYYYMMDD, time.Now().Format(utils.LayoutYYYYMMDD))
	for _, d := range v.Daily {
		dd, _ := time.Parse(utils.LayoutYYYYMMDD, d.Date)
		if dd.Sub(nowDate) < time.Duration(limit)*day {
			daily = append(daily, d)
		}
	}
	v.Daily = daily
	return v
}

func (v *WeatherIndicesResp) FilterType(it string) *WeatherIndicesResp {
	if it == "0" || it == "" {
		return v
	}
	types := strings.Split(it, ",")
	var daily = make([]*WeatherIndicesRespDaily, 0)
	for _, d := range v.Daily {
		if utils.ContainStr(types, d.Type) {
			daily = append(daily, d)
		}
	}
	v.Daily = daily
	return v
}

type RealtimeAirRespNow struct {
	PubTime  string `json:"pubTime"`
	Aqi      string `json:"aqi"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Primary  string `json:"primary"`
	Pm10     string `json:"pm10"`
	Pm2p5    string `json:"pm2p5"`
	No2      string `json:"no2"`
	So2      string `json:"so2"`
	Co       string `json:"co"`
	O3       string `json:"o3"`
}

type RealtimeAirRespStation struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	PubTime  string `json:"pubTime"`
	Aqi      string `json:"aqi"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Primary  string `json:"primary"`
	Pm10     string `json:"pm10"`
	Pm2p5    string `json:"pm2p5"`
	No2      string `json:"no2"`
	So2      string `json:"so2"`
	Co       string `json:"co"`
	O3       string `json:"o3"`
}

type RealtimeAirResp struct {
	Code       string                    `json:"code"`
	UpdateTime string                    `json:"updateTime"`
	FxLink     string                    `json:"fxLink"`
	Now        *RealtimeAirRespNow       `json:"now"`
	Station    []*RealtimeAirRespStation `json:"station"`
}

func (v *RealtimeAirResp) GetCode() string {
	return v.Code
}

func (v *RealtimeAirResp) GetType() enum.WeatherInfoType {
	return enum.WeatherInfo_AIR_REALTIME
}

func (v *RealtimeAirResp) String() string {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(bytes)
}

func (v *RealtimeAirResp) GetUpdateTime() (time.Time, error) {
	return utils.ParseTimeFromIOS8601(v.UpdateTime)
}

func (v *RealtimeAirResp) FormatTime() QWResp {
	v.UpdateTime = utils.ConvertTimeFromIOS8601(v.UpdateTime)
	v.Now.PubTime = utils.ConvertTimeFromIOS8601(v.Now.PubTime)
	for _, h := range v.Station {
		h.PubTime = utils.ConvertTimeFromIOS8601(h.PubTime)
	}
	return v
}

func (v *RealtimeAirResp) Limit(limit uint) QWResp {
	return v
}