package db

import (
	"fmt"
	"time"

	"github.com/MarcelCode/ROWA/src/util"
	_ "github.com/mattn/go-sqlite3"
)

type Times struct {
	TimeOn  string `json:"time_on"`
	TimeOff string `json:"time_off"`
}

type CurrentTime struct {
	TimeOn  string `json:"time_on"`
	TimeOff string `json:"time_off"`
	State   int    `json:"state"`
}

type PumpData struct {
	TimeOn   string `json:"time_on"`
	Duration int    `json:"duration"`
}

type PlantTypes struct {
	TypeName   string `json:"typename"`
	TypeModule int    `json:"typemodule"`
}



type RealityCheckData struct  {
	ModuleNumber int `json:"modulenum"`
	Age []int `json:"age"`
	Type string `json:"type"`
}

type RealityCheckDataPos struct {
	Age int `json:"age"`
	Harvestable bool `json:"harvestable"`
}

type LightState struct {
	State int `json:"state"`
}

type KnownType struct {
	KnownType string `json:"type"`
}

//store incoming new time data

//get time and state from db
func (store *Database) GetLightTimes() (currentTime *CurrentTime, err error) {

	sqlQuery := `SELECT OnTime, OffTime, CurrentState
				 FROM TimeTable
				 WHERE ID = 1`

	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}

	currentTime = &CurrentTime{}
	rows.Next()
	err = rows.Scan(&currentTime.TimeOn, &currentTime.TimeOff, &currentTime.State)

	if err != nil {
		return
	}
	//get Times and State from DB
	//make Json and send

	return
}

func (store *Database) InsertLightTimes(newTimes *Times) (status *Status, err error) {

	status = &Status{}

	sqlQuery := `UPDATE TimeTable SET OnTime = ?, OffTime = ? WHERE ID = 1`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()

	_, err = statement.Exec(newTimes.TimeOn, newTimes.TimeOff)
	if err != nil {
		status.Message = "error"
		return
	}
	util.LightTimesRenew()

	status.Message = "Light Inserted"
	return
}

func (store *Database) GetPlantTypes() (plantPossible []*PlantTypes, err error) {

	sqlQuery := `SELECT PlantType, Position
				 FROM Module`

	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		plantTypes := &PlantTypes{}
		err = rows.Scan(&plantTypes.TypeName, &plantTypes.TypeModule)

		if err != nil {
			return
		}
		plantPossible = append(plantPossible, plantTypes)
	}
	return
}

func (store *Database) InsertModuleChanges(plantTypes *PlantTypes) (status *Status, err error) {

	status = &Status{}

	sqlQuery := `UPDATE Module SET PlantType = ?, AvailableSpots = TotalSpots WHERE Position = ?`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()
	_, err = statement.Exec(plantTypes.TypeName, plantTypes.TypeModule)
	
	if err != nil {
		status.Message = "error"
		return
	}

	sqlQuery = `UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = ? AND Module= ?`
	statement, _ = store.Db.Prepare(sqlQuery)
	for i := 0; i < 6; i++ {
	_, err = statement.Exec(i+1, plantTypes.TypeModule)
	
	}
	status.Message = "Module Changed"
	return
}

func (store *Database) GetKnownPlantTypes() (knownTypes []*KnownType, err error) {

	sqlQuery := `SELECT Name
	FROM PlantType`

	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		knownType := &KnownType{}
		err = rows.Scan(&knownType.KnownType)

		if err != nil {
			fmt.Println(err)
			return

		}
		knownTypes = append(knownTypes, knownType)
	}

	return
}

func (store *Database) GetPumpTime() (pumpData *PumpData, err error) {

	sqlQuery := `SELECT OnTime, CurrentState
				 FROM TimeTable
				 WHERE ID = 2`

	rows, err := store.Db.Query(sqlQuery)
	defer rows.Close()
	if err != nil {
		return
	}

	pumpData = &PumpData{}
	rows.Next()
	err = rows.Scan(&pumpData.TimeOn, &pumpData.Duration)

	if err != nil {
		return
	}

	return
}

func (store *Database) InsertPumpTimes(pumpData *PumpData) (status *Status, err error) {

	status = &Status{}

	sqlQuery := `UPDATE TimeTable SET OnTime = ?, CurrentState = ? WHERE ID = 2`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()
	_, err = statement.Exec(pumpData.TimeOn, pumpData.Duration)
	if err != nil {
		status.Message = "error"
		return
	}
	util.PumpTimesRenew()

	status.Message = "Pump Time Inserted"
	return
}

func (store *Database) RealityCheck(realitycheckData *RealityCheckData)  (status *Status, err error) {
	status = &Status{}

	//Insert RealityChecked Data into tables
	for i, v := range realitycheckData.Age {
		if (v == 0){
			sqlQuery := `UPDATE Plant SET Harvested = 1, PlantPosition = 0 WHERE PlantPosition = ? AND Module= ?`
			statement, _ := store.Db.Prepare(sqlQuery)
			defer statement.Close()

			_, err = statement.Exec(i+1, realitycheckData.ModuleNumber)
			if err != nil {
				status.Message = "error"
				return
			}
			sqlQuery = `UPDATE Module SET AvailableSpots = AvailableSpots + 1 WHERE Position= ?`
			statement, _ = store.Db.Prepare(sqlQuery)

			_, err = statement.Exec(realitycheckData.ModuleNumber)
			if err != nil {
				status.Message = "error"
				return
			}
			
		} else if (v == 1){
			sqlQuery := `DELETE FROM Plant WHERE PlantPosition = ? AND Module= ? AND Harvested = 0`
			statement, _ := store.Db.Prepare(sqlQuery)
			defer statement.Close()
			_, err = statement.Exec(i+1, realitycheckData.ModuleNumber)

			sqlQuery = `INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)`
			statement, _ = store.Db.Prepare(sqlQuery)
			defer statement.Close()
			_, err = statement.Exec(realitycheckData.ModuleNumber, i+1, time.Now().Format("2006-01-02"), 0)
           
		} else{
			sqlQuery := `DELETE FROM Plant WHERE PlantPosition = ? AND Module= ? AND Harvested = 0`
			statement, _ := store.Db.Prepare(sqlQuery)
			defer statement.Close()
			_, err = statement.Exec( i+1, realitycheckData.ModuleNumber)

			sqlQuery = `INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)`
			statement, _ = store.Db.Prepare(sqlQuery)
			defer statement.Close()
			_, err = statement.Exec(realitycheckData.ModuleNumber, i+1, time.Now().AddDate(0,0,-v).Format("2006-01-02"), 0)
		
					
		}
	}


	//Update Available Plant Count
	sqlQuery := `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
	rows, err := store.Db.Query(sqlQuery, realitycheckData.ModuleNumber)
	rows.Next()
	var id int

	rows.Scan(&id)
	rows.Close()

	sqlQuery = `UPDATE Module SET AvailableSpots = TotalSpots - ? WHERE Position = ?`
	_, err = store.Db.Exec(sqlQuery, id, realitycheckData.ModuleNumber)

	if err != nil {
		status.Message = "error"
		return
	}

	status.Message = "RealityCheck Data inserted"
	return 

} 
