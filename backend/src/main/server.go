package main

import (
	"database/sql"
	"os"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"

	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/MarcelCode/ROWA/src/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var filename string = "logfile.log"
    // Create the log file if doesn't exist. And append to it if it already exists.
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
    Formatter := new(log.TextFormatter)
    // You can change the Timestamp format. But you have to use the same date and time.
    // "2006-02-02 15:04:06" Works. If you change any digit, it won't work
    // ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
    log.SetFormatter(Formatter)
    if err != nil {
        // Cannot open log file. Logging to stderr
        fmt.Println(err)
    }else{
        log.SetOutput(f)
	}
	

	database, err := sql.Open("sqlite3", "rowa.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	log.Warning("This is a warning")
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
