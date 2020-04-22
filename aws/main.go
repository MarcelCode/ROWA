package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type SensorValues struct {
	FarmId         string  `json:"farm_id"`
	Temp           float32 `json:"temp"`
	LightIntensity float32 `json:"lightIntensity"`
	Humid          float32 `json:"humidity"`
	WaterLevel     float32 `json:"waterLevel"`
	WaterTemp      float32 `json:"waterTemp"`
	WaterpH        float32 `json:"waterpH"`
	Timestamp      int64   `json:"timestamp"`
}

func HandleRequest(ctx context.Context, values SensorValues) {
	client := influxdb2.NewClient("https://eu-central-1-1.aws.cloud2.influxdata.com", os.Getenv("TestValueToken"))
	//non-blocking write client
	log.Printf("%+v\n", values)
	log.Println(values.Temp)
	writeApi := client.WriteApi(os.Getenv("OrgID"), os.Getenv("Bucket"))
	data := influxdb2.NewPointWithMeasurement("sensor-data").
		AddTag("farm-id", values.FarmId).
		AddField("temp", values.Temp).
		AddField("lightIntensity", values.LightIntensity).
		AddField("humidity", values.Humid).
		AddField("waterLevel", values.WaterLevel).
		AddField("waterTemp", values.WaterTemp).
		AddField("waterpH", values.WaterpH).
		SetTime(time.Unix(values.Timestamp, 0))
	// write point asynchronously
	writeApi.WritePoint(data)
	// Flush writes
	writeApi.Flush()
	defer client.Close()
}

func main() {
	lambda.Start(HandleRequest)
}
