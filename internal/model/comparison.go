package model

type ComparisonResult struct {
	Season  string `json:"season"`
	DriverA string `json:"driver_a"`
	DriverB string `json:"driver_b"`
	WinsA   int    `json:"wins_a"`
	WinsB   int    `json:"wins_b"`
}
