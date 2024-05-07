package main

import (
	"encoding/json"
)

type observationData struct {
	BirdClassId int     `json:"class_id"`
	Confidence  float64 `json:"confidence"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	Timestamp   float64 `json:"timestamp"`
}

func readJsonFile(data []byte) []observationData {
	var fileData []observationData
	json.Unmarshal(data, &fileData)
	return fileData
}
