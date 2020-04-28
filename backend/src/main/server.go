package main

import (
	"database/sql"
	
	log "github.com/sirupsen/logrus"

	"os"
	"time"
	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"

	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/MarcelCode/ROWA/src/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)



func main() {
	
	util.LogNameIntervall()
    go util.RunnerLog()

	database, err := sql.Open("sqlite3", "rowa.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	
	db.InitStore(&db.Database{Db: database})

	if settings.Debug {
		db.FunctionStore.DbSetup()
		

	}

	if settings.ArduinoOn {
		go sensor.ReadSensorData()
		util.LightTimesRenew()
		util.PumpTimesRenew()
		go util.Runner()
	} else {
		go sensor.ReadFakeSensorData()
	}

	e := echo.New()
	e.Use(middleware.CORS())


	//make echo log all requests
	var filename string = "logfile-httprequests-"+ time.Now().Format("2006-01-02") +".log"
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
		Output: f,
	  }))

	// Routes
	e.GET("/dashboard/sensor-data", api.GetSensorDataHandler)
	e.GET("/dashboard/harvestable-plants", api.GetHarvestablePlantsHandler)
	e.GET("/dashboard/plantable-plants", api.GetPlantablePlantsHandler)

	e.GET("/harvest/get-plant", api.GetHarvestablePlantHandler)
	e.POST("/harvest/harvestdone", api.HarvestDoneHandler)

	e.GET("/plant/blinkstop", api.StopModuleBlink)
	e.GET("/plant/get-position", api.PlantHandler)
	e.POST("/plant/finish", api.FinishPlantingHandler)
	e.GET("/dashboard/cattree/:module", api.GetCatTreeDataHandler)

	e.POST("/adminSettings/insert-light", api.InsertLightTimes)
	e.GET("/adminSettings/get-light", api.GetLightTimes)
	e.POST("/adminSettings/changelight", api.ChangeLightState)

	e.GET("/adminSettings/get-types", api.GetPlantTypes)
	e.GET("/adminSettings/get-knowntypes", api.GetKnownPlantTypes)
	e.POST("/adminSettings/insertmodule-change", api.InsertModuleChanges)

	e.POST("/adminSettings/insert-pump", api.InsertPumpTime)
	e.GET("/adminSettings/get-pump", api.GetPumpTimes)

	e.GET("/plant/get-all", api.AllPlantHandler)
	e.POST("/plant/plant-all", api.MassPlantingHandler)

	e.GET("/harvest/get-all", api.AllHarvestHandler)
	e.POST("/harvest/harvest-all", api.MassHarvestHandler)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))

}
