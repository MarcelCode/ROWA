package influx

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type State struct {
	State *Reported `json:"state"`
}
type Reported struct {
	Reported *SensorData `json:"reported"`
}
type SensorData struct {
	FarmId         string  `json:"farm_id"`
	Temp           float32 `json:"temp"`
	LightIntensity float32 `json:"lightIntensity"`
	Humid          float32 `json:"humidity"`
	WaterLevel     float32 `json:"waterLevel"`
	WaterTemp      float32 `json:"waterTemp"`
	WaterpH        float32 `json:"waterpH"`
	Timestamp      int64   `json:"timestamp"`
}

func AwsInit() *iotdataplane.IoTDataPlane {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"), Endpoint: aws.String("arig5vtggzkh-ats.iot.eu-central-1.amazonaws.com")},
	)
	if err != nil {
		log.Fatal(err)
	}
	iotDataSvc := iotdataplane.New(sess)
	return iotDataSvc
}
func AwsPublishInput(iotDataSvc *iotdataplane.IoTDataPlane, s []float32, timestamp int64) {
	data := &State{&Reported{&SensorData{
		FarmId:         "1",
		Temp:           s[0],
		LightIntensity: s[1],
		Humid:          s[2],
		WaterLevel:     s[3],
		WaterTemp:      s[4],
		WaterpH:        s[5],
		Timestamp:      timestamp,
	}}}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("could not marshal JSON: %s", err)
	}
	fmt.Println(string(jsonStr))
	input := &iotdataplane.PublishInput{
		Payload: jsonStr,
		Topic:   aws.String("$aws/things/farm_1/shadow/update"),
		Qos:     aws.Int64(0),
	}
	resp, err := iotDataSvc.Publish(input)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
