package db

import (
	"database/sql"
	"log"
	"time"
)

type SensorData struct {
	Datetime       string  `json:"datetime"`
	Temp           float64 `json:"temperature"`
	LightIntensity float64 `json:"light_intensity"`
	Humidity       float64 `json:"humidity"`
	WaterLevel     float64 `json:"water_level"`
	WaterTemp      float64 `json:"water_temp"`
	WaterpH        float64 `json:"water_ph"`
}

type PlantsPerPlantType struct {
	Name            string `json:"plant_type"`
	AvailablePlants int    `json:"available_plants"`
}

type BlinkModule struct {
	Module int `json:"module"`
}

type PlantInfoPerModule struct {
	PlantType     string `json:"plant_type"`
	PlantPosition int    `json:"plant_position"`
	Age           int    `json:"age"`
	Harvestable   bool   `json:"harvestable"`
}

func (store *Database) GetHarvestablePlants() (plantsToHarvest []*PlantsPerPlantType, err error) {
	sqlQuery := `SELECT PlantType, COUNT(PlantType) as AvailablePlantsPerPlantType
				FROM Plant
						 INNER JOIN Module M on Plant.Module = M.Id
						 INNER JOIN PlantType PT on M.PlantType = PT.Name
				where Harvested = 0
				  and date(PlantDate, '+' || GrowthTime || ' days') <= date('now')
				GROUP BY PlantType`
	rows, err := store.Db.Query(sqlQuery)
	HandleError(err)
	defer rows.Close()

	for rows.Next() {
		plantsPerPlantType := &PlantsPerPlantType{}
		err = rows.Scan(&plantsPerPlantType.Name, &plantsPerPlantType.AvailablePlants)
		if err != nil {
			log.Fatal(err)
		}
		plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)
	}
	return
}

func (store *Database) GetPlantablePlantsPerModule(rows *sql.Rows, moduleNumber int) (plantsToPlant []*PlantsPerPlantType) {
	plantsPerModule := store.GetAmountOfPlantsPerModule(moduleNumber)
	for rows.Next() {
		plantsPerPlantType := &PlantsPerPlantType{}
		hasAvailableSpots := plantsPerPlantType.AvailablePlants-plantsPerModule == 0 && plantsPerModule < 6

		err := rows.Scan(&plantsPerPlantType.Name, &plantsPerPlantType.AvailablePlants)

		if hasAvailableSpots {
			plantsToPlant = matchPlantsToType(plantsPerPlantType, plantsToPlant)
		}
		if err != nil && plantsPerModule < 6 {
			plantsPerPlantType = &PlantsPerPlantType{store.GetPlantTypePerModule(moduleNumber), 6}
			plantsToPlant = matchPlantsToType(plantsPerPlantType, plantsToPlant)
			break
		}

	}
	return plantsToPlant
}

func (store *Database) GetAllPlantablePlants() (plantablePlants []*PlantsPerPlantType, err error) {

	sqlQuery := ``

	for module := 1; module < 7; module++ {
		sqlQuery = `SELECT PlantType, COUNT(PlantType) as AvailablePlants
					FROM Plant
							INNER JOIN Module M on Plant.Module = M.Id
							INNER JOIN PlantType PT on M.PlantType = PT.Name
					WHERE Harvested = 0
					AND M.Id = ?
					AND date(PlantDate, '+' || 7 || ' days') <= date('now')`

		rows, err := store.Db.Query(sqlQuery, module)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		plantablePlants = store.GetPlantablePlantsPerModule(rows, module)

	}
	return plantablePlants, err
}

func (store *Database) GetAllPlantsInModules() (plantsToHarvest []*PlantsPerPlantType, err error) {
	sqlQuery := ``
	for i := 1; i < 7; i++ {
		sqlQuery = `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
		rows, err := store.Db.Query(sqlQuery, i)
		rows.Next()
		var id int
		rows.Scan(&id)

		if err != nil {
			log.Fatal(err)
		}
		rows.Close()

		sqlQuery = `SELECT PlantType, COUNT(PlantType) as AvailablePlants
					FROM Plant
							INNER JOIN Module M on Plant.Module = M.Id
							INNER JOIN PlantType PT on M.PlantType = PT.Name
					WHERE Harvested = 0
					AND M.Id = ?
					AND date(PlantDate, '+' || 7 || ' days') <= date('now')`

		rows, err = store.Db.Query(sqlQuery, i)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {

			plantsPerPlantType := &PlantsPerPlantType{}
			err = rows.Scan(&plantsPerPlantType.Name, &plantsPerPlantType.AvailablePlants)
			if plantsPerPlantType.AvailablePlants-id == 0 && 6-id > 0 {
				plantsPerPlantType.AvailablePlants = i
				if err != nil {
					sqlQuery = `SELECT PlantType
				FROM Module 
				WHERE Id = ?`
					rows, err = store.Db.Query(sqlQuery, i)
					var plantType string
					rows.Scan(&plantType)
					plantsPerPlantType := &PlantsPerPlantType{}
					plantsPerPlantType.Name = plantType
					plantsPerPlantType.AvailablePlants = 6
					if plantsPerPlantType.AvailablePlants-id == 0 && 6-id > 0 {
						plantsPerPlantType.AvailablePlants = i
						if err != nil {
							log.Print(err)
						}
						plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)
					}
					break
				}
				plantsToHarvest = append(plantsToHarvest, plantsPerPlantType)
			}
			break

		}
	}
	return

}
func (store *Database) GetAmountOfPlantsPerModule(moduleNumber int) (id int) {
	sqlQuery := `SELECT COUNT(Id) FROM Plant WHERE Harvested = 0 AND Module = ?`
	rows, err := store.Db.Query(sqlQuery, moduleNumber)
	rows.Next()

	rows.Scan(&moduleNumber)

	if err != nil {
		log.Fatal(err)
	}
	rows.Close()
	return moduleNumber
}

func (store *Database) GetPlantTypePerModule(moduleNumber int) (plantType string) {
	sqlQuery := `SELECT PlantType
					FROM Module 
					WHERE Id = ?`
	rows, err := store.Db.Query(sqlQuery, moduleNumber)
	rows.Next()
	rows.Scan(&plantType)
	if err != nil {
		log.Print(err)
	}
	return plantType
}

func matchPlantsToType(plantsPerPlantType *PlantsPerPlantType, plantablePlants []*PlantsPerPlantType) []*PlantsPerPlantType {

	p, found := find(plantablePlants, plantsPerPlantType.Name)
	if found {
		plantablePlants[p].AvailablePlants++
	} else {
		plantsPerPlantType.AvailablePlants = 1
		plantablePlants = append(plantablePlants, plantsPerPlantType)
	}
	return plantablePlants
}

func (store *Database) GetLastSensorEntry() (sensorData *SensorData, err error) {
	sqlQuery := `SELECT Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH
				 FROM SensorMeasurements
				 WHERE ID = (SELECT MAX(ID)  FROM SensorMeasurements)`

	row, err := store.Db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	sensorData = &SensorData{}
	row.Next()
	err = row.Scan(&sensorData.Datetime, &sensorData.Temp, &sensorData.LightIntensity, &sensorData.Humidity, &sensorData.WaterLevel, &sensorData.WaterTemp, &sensorData.WaterpH)

	if err != nil {
		log.Fatal(err)
	}

	return sensorData, err
}

func (store *Database) GetCatTreeData(module int) (plantInfo []*PlantInfoPerModule, err error) {
	sqlQuery := `SELECT PlantType, PlantPosition, PlantDate, GrowthTime
				FROM Module M
				INNER JOIN Plant P ON M.Position = P.Module
				INNER JOIN PlantType PT ON M.PlantType = PT.Name
				WHERE M.Position = ? AND Harvested = 0
				ORDER BY PlantPosition`
	rows, err := store.Db.Query(sqlQuery, module)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		plantInfoPerModule := &PlantInfoPerModule{}
		var plantDate = ""
		var growthTime = 0

		err = rows.Scan(&plantInfoPerModule.PlantType, &plantInfoPerModule.PlantPosition, &plantDate, &growthTime)
		if err != nil {
			log.Fatal(err)
		}

		layout := "2006-01-02"
		plantDateParsed, err := time.Parse(layout, plantDate)
		if err != nil {
			log.Fatal(err)
		}
		plantInfoPerModule.Age = int(time.Now().Sub(plantDateParsed).Hours() / 24)
		if plantInfoPerModule.Age >= growthTime {
			plantInfoPerModule.Harvestable = true
		} else {
			plantInfoPerModule.Harvestable = false
		}
		plantInfo = append(plantInfo, plantInfoPerModule)
	}

	return plantInfo, err

}

func find(slice []*PlantsPerPlantType, val string) (int, bool) {
	for i, item := range slice {
		if item.Name == val {
			return i, true
		}
	}
	return -1, false
}
