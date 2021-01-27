package api

import (
	"net/http"
	"strconv"

	"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/labstack/echo"
	"github.com/MarcelCode/ROWA/src/settings"
)

func GetHarvestablePlantsHandler(c echo.Context) (err error) {
	plantsToHarvest, err := db.FunctionStore.GetHarvestablePlants()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Harvestable Plants not found")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func GetPlantablePlantsHandler(c echo.Context) (err error) {
	plantsToHarvest, err := db.FunctionStore.GetPlantablePlants()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Plantable Plants not found")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func GetPlantableModulesHandler(c echo.Context) (err error) {
	plantsToHarvest, err := db.FunctionStore.GetAllPlantsInModules()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Plantable Plants not found")
	}

	return c.JSON(http.StatusOK, plantsToHarvest)
}

func GetSensorDataHandler(c echo.Context) (err error) {
	sensorData, err := db.FunctionStore.GetLastSensorEntry()

	if err != nil {
		return c.JSON(http.StatusNotFound, "Sensor Data not found")
	}

	return c.JSON(http.StatusOK, sensorData)
}

func GetCatTreeDataHandler(c echo.Context) (err error) {
	module, err := strconv.Atoi(c.Param("module")) // add err handling
	catTreeObject, err := db.FunctionStore.GetCatTreeData(module)

	if err != nil {
		return c.JSON(http.StatusNotFound, "Cat Tree Data not found")
	}

	return c.JSON(http.StatusOK, catTreeObject)

}

func StartBlink(c echo.Context) (err error) {
	 
	blinkModule := &db.BlinkModule{}
	c.Bind(&blinkModule)
	
	if settings.ArduinoOn {
		go sensor.ActivateModuleLight(blinkModule.Module)
	}

	return c.JSON(http.StatusOK, "Light Triggered")
}

