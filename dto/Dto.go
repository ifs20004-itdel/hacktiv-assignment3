package dto

type DTO struct {
	WaterStatus string `json:"waterStatus"`
	WindStatus  string `json:"windStatus"`
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
}
