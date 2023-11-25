package entity

type MapElement struct {
	Id        int32    `json:"id"`
	Type      *string  `json:"type"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Title     *string  `json:"title"`
}
