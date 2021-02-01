package db

import "github.com/labstack/gommon/log"

type PlantType struct {
	Name string `json:"plant_type" query:"plant_type"`
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
