package db

import (
	"fmt"

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

	sqlQuery := `UPDATE Module SET PlantType = ? WHERE Position = ?`
	statement, _ := store.Db.Prepare(sqlQuery)
	defer statement.Close()

	_, err = statement.Exec(plantTypes.TypeName, plantTypes.TypeModule)
	if err != nil {
		status.Message = "error"
		return
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
