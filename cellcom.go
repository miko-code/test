package main

import "time"

// Cellcom data struct
type CellcomSlice struct {
	Cellcome []struct {
		Time       time.Time `json:"Time"`
		Latitude   int       `json:"Latitude"`
		Longitude  int       `json:"Longitude"`
		Value      int       `json:"Value"`
		DeviceName string    `json:"DeviceName"`
		Type       string    `json:"Type"`
	} `json:"data"`
}
