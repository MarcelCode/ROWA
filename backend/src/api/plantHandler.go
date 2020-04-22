package api

import (
	"fmt"
	"net/http"

	"github.com/MarcelCode/ROWA/src/db"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/echo"
)

func PlantHandler(c echo.Context) (err error) {
	plantType := &db.PlantType{}
	err = c.Bind(plantType)
	position, err := db.FunctionStore.Plant(plantType)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Planting not available")
	}

	return c.JSON(http.StatusOK, position)
}

func FinishPlantingHandler(c echo.Context) (err error) {
	plantedModule := &db.PlantedModule{}
	c.Bind(&plantedModule)

	status, err := db.FunctionStore.FinishPlanting(plantedModule)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, status)
}

func AllPlantHandler(c echo.Context) (err error) {
	position, err := db.FunctionStore.AllPlantable()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Planting not available")
	}

	return c.JSON(http.StatusOK, position)
}

func MassPlantingHandler(c echo.Context) (err error) {
	fmt.Println("Server")
	plantedModules := &db.PlantedModules{}
	c.Bind(plantedModules)
	fmt.Println(c)
	status, err := db.FunctionStore.MassPlanting(plantedModules)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, status)
}

func StopModuleBlink(c echo.Context) (err error) {
	if settings.ArduinoOn {
		sensor.DeactivateModuleLight()
	}
	if err != nil {
		return c.JSON(http.StatusNotFound, "Finish planting not possible")
	}

	return c.JSON(http.StatusOK, "Canceled Light")
}
