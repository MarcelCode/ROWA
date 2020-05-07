package main

import (
	"database/sql"
	
	log "github.com/sirupsen/logrus"
	"fmt"
	"time"
	"github.com/MarcelCode/ROWA/src/api"
	"github.com/MarcelCode/ROWA/src/db"

	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/MarcelCode/ROWA/src/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)



func main() {

	
	
	InitBackendLogger()
	

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





	logf, err := rotatelogs.New(
		"httprequest.%Y%m%d%H%M%S.log",
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour*3),
	  )
	  if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	  }

	//make echo log all requests
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
		Output: logf,
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


func InitBackendLogger(){

	logb, err := rotatelogs.New(
		"backend.%Y%m%d%H%M%S.log",
		rotatelogs.WithMaxAge(24 * time.Hour),
		rotatelogs.WithRotationTime(time.Hour*3),
	  )
	  if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	  }

	Formatter := new(log.TextFormatter)
    // You can change the Timestamp format. But you have to use the same date and time.
    // "2006-02-02 15:04:06" Works. If you change any digit, it won't work
    // ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
     Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	
	log.SetFormatter(Formatter)
    if err != nil {
        // Cannot open log file. Logging to stderr
        log.Println(err)
    }else{
        log.SetOutput(logb)
	}

}

func MoveFile(Location string){
	fmt.Println(Location)
}


