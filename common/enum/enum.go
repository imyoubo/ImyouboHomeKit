package enum

type WeatherInfoType byte

const (
	_ WeatherInfoType = iota
	WeatherInfo_WEATHER_REALTIME
	WeatherInfo_WEATHER_15DAYS
	WeatherInfo_WEATHER_24HOURS
	WeatherInfo_AIR_REALTIME
	WeatherInfo_INDICES_1DAYS
)

func (wt WeatherInfoType) String() string {
	switch wt {
	case WeatherInfo_WEATHER_REALTIME:
		return "WEATHER_REALTIME"
	case WeatherInfo_WEATHER_15DAYS:
		return "WEATHER_15DAYS"
	case WeatherInfo_WEATHER_24HOURS:
		return "WEATHER_24HOURS"
	case WeatherInfo_AIR_REALTIME:
		return "AIR_REALTIME"
	case WeatherInfo_INDICES_1DAYS:
		return "INDICES_1DAYS"
	default:
		return "UNKNOWN"
	}
}