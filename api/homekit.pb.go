package api

type ListLocationInfoRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Kw     string `json:"kw"`
}

type GetRealtimeWeather struct {
	Location uint64 `json:"location"`
	Refresh  bool   `json:"refresh"`
}
