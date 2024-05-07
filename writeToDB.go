package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

var c = 0

func writeData(data observationData, configData *config) {
	token := configData.InfluxDBToken
	url := configData.InfluxDBUrl
	client := influxdb2.NewClient(url, token)

	org := configData.InfluxDBOrga
	bucket := configData.InfluxDBBucket
	writeAPI := client.WriteAPIBlocking(org, bucket)

	tags := map[string]string{
		"birdId":    fmt.Sprintf("%d", data.BirdClassId),
		"latitude":  fmt.Sprintf("%f", data.Latitude),
		"longitude": fmt.Sprintf("%f", data.Longitude),
		//"mID":       fmt.Sprintf("%d", c),
	}
	fields := map[string]interface{}{
		"confidence": data.Confidence,
	}

	dataPoint := write.NewPoint("BirdObservation", tags, fields, time.Unix(int64(data.Timestamp), 0))

	if err := writeAPI.WritePoint(context.Background(), dataPoint); err != nil {
		fmt.Println(data)
		fmt.Println(err)
	} else {
		fmt.Printf("Data written successfully %d %s\n", c, dataPoint.Time())
	}
	c++
}
