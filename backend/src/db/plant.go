package db

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/MarcelCode/ROWA/src/settings"
	"github.com/labstack/gommon/log"
)

type PlantedModule struct {
	Module int `json:"planted_module" query:"planted_module"`
}

type PlantedModules struct {
	Module string `json:"planted_modules" query:"planted_modules"`
}
type PlantableModules struct {
	PlantType string `json:"plant_type" query:"plant_type"`
	Module    int    `json:"planted_module" query:"planted_module"`
}

func (store *Database) Plant(plantType *PlantType) (modulePosition int, err error) {
	sqlQuery := `SELECT Position FROM Module WHERE AvailableSpots > 0 AND PlantType= ?`
	rows, err := store.Db.Query(sqlQuery, plantType.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&modulePosition)
		if err != nil {
			return
		}
	} else {
		err = errors.New("no data available")
		return
	}

	if settings.ArduinoOn {
		go sensor.ActivateModuleLight(modulePosition)
	}

	return
}
func (store *Database) AllPlantable() (plantableModules []*PlantableModules, err error) {
	sqlQuery := `SELECT Position, PlantType FROM Module WHERE AvailableSpots > 0`
	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		plantableModule := &PlantableModules{}
		err = rows.Scan(&plantableModule.Module, &plantableModule.PlantType)

		if err != nil {
			fmt.Println(err)
			return

		}
		plantableModules = append(plantableModules, plantableModule)

		if settings.ArduinoOn {
			go sensor.ActivateModuleLight(plantableModule.Module)
		}
	}

	

	return
}

func (store *Database) FinishPlanting(plantedModule *PlantedModule) (status *Status, err error) {
	sqlQuery := `UPDATE Plant SET PlantPosition = PlantPosition + 1 WHERE Harvested = 0 AND Module = ?`
	_, err = store.Db.Exec(sqlQuery, plantedModule.Module)

	sqlQuery = `INSERT INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)`
	statement, _ := store.Db.Prepare(sqlQuery)
	_, err = statement.Exec(plantedModule.Module, 1, time.Now().Format("2006-01-02"), 0)

	sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
	rows, err := store.Db.Query(sqlQuery, plantedModule.Module)
	rows.Next()
	var id int

	rows.Scan(&id)
	rows.Close()

	sqlQuery = `UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?`
	_, err = store.Db.Exec(sqlQuery, id, plantedModule.Module)
	fmt.Println(err)

	if settings.ArduinoOn {
		go sensor.DeactivateModuleLight()
	}

	status = &Status{Message: "Planting Done"}
	return
}

func (store *Database) MassPlanting(plantedModules *PlantedModules) (status *Status, err error) {
	data_array := strings.Split(plantedModules.Module, ",")

	for _, plantedModule := range data_array {
		plantedModuleInt, _ := strconv.Atoi(plantedModule)
	/*	var EmptySpots int
		sqlQuery := `SELECT AvailableSpots
		FROM Module
		WHERE ID = ?`
		rows, err := store.Db.Query(sqlQuery, plantedModuleInt)
		rows.Next()
		err = rows.Scan(&EmptySpots)

		rows.Close()
		fmt.Println(EmptySpots)
		fmt.Println("")
		for i := 0; i < EmptySpots; i++{*/  //Code if all Plantables in each given Module should be planted(:= declarations need to be chaged to = if this code is added again)
		sqlQuery := `UPDATE Plant SET PlantPosition = PlantPosition + 1 WHERE Harvested = 0 AND Module = ?`
		_, err = store.Db.Exec(sqlQuery, plantedModuleInt)
		fmt.Println(plantedModule)
		fmt.Println("")
		sqlQuery = `INSERT INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)`
		statement, _ := store.Db.Prepare(sqlQuery)
		_, err = statement.Exec(plantedModuleInt, 1, time.Now().Format("2006-01-02"), 0)

		sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
		rows, err := store.Db.Query(sqlQuery, plantedModuleInt)
		rows.Next()
		var id int

		rows.Scan(&id)
		rows.Close()

		sqlQuery = `UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?`
		_, err = store.Db.Exec(sqlQuery, id, plantedModuleInt)
		fmt.Println(err)
	//}
		
	}
	if settings.ArduinoOn {
		go sensor.DeactivateModuleLight()
	}
	status = &Status{Message: "Planting Done"}
	return
}
