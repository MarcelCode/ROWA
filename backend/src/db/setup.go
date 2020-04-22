package db

import (
	_ "database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func (store *Database) DbSetup() (err error) {
	//Creating PlantType DB
	statement, _ := store.Db.Prepare("CREATE TABLE IF NOT EXISTS PlantType (Name TEXT PRIMARY KEY, GrowthTime INTEGER)") // TODO Why no ID?
	statement.Exec()
	statement, _ = store.Db.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", 42) //TODO Why only one PlantType, whether we use two plants?
	statement.Exec("Basil", 21)
	statement.Exec("Mache", 32)
	statement.Exec("Spinach", 30)
	rows, _ := store.Db.Query("SELECT Name, Growthtime from PlantType")
	defer rows.Close()
	var name string
	var growthTime int
	for rows.Next() {
		rows.Scan(&name, &growthTime)
		fmt.Println(name + ": " + strconv.Itoa(growthTime))
	}
	//Creating Module DB
	statement, _ = store.Db.Prepare("DROP TABLE IF EXISTS Module")
	statement.Exec()
	statement, _ = store.Db.Prepare("CREATE TABLE IF NOT EXISTS Module (Id INTEGER PRIMARY KEY, PlantType TEXT, Position INTEGER UNIQUE, AvailableSpots INTEGER, TotalSpots INTEGER, FOREIGN KEY(PlantType) REFERENCES PlantType(Name))")
	statement.Exec()
	statement, _ = store.Db.Prepare("INSERT OR IGNORE INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec(1, "Basil", 0, 6)
	statement.Exec(2, "Mache", 2, 6)
	statement.Exec(3, "Lettuce", 0, 6)
	statement.Exec(4, "Mache", 1, 6)
	statement.Exec(5, "Lettuce", 0, 6)
	statement.Exec(6, "Basil", 0, 6)

	// Creating Plant DB
	statement, _ = store.Db.Prepare("DROP TABLE IF EXISTS Plant")
	statement.Exec()
	statement, _ = store.Db.Prepare("CREATE TABLE IF NOT EXISTS Plant (Id INTEGER PRIMARY KEY, Module INTEGER, PlantPosition INTEGER, PlantDate TEXT, Harvested INTEGER, FOREIGN KEY (Module) REFERENCES Module(Id))")
	statement.Exec()
	statement, _ = store.Db.Prepare("INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")

	// Add multiple Plants for testing purposes
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			statement.Exec(i+1, j+1, time.Now().Format("2006-01-02"), 0)
		}
	}

	statement, _ = store.Db.Prepare("UPDATE Plant SET PlantDate = '2019-09-09'  WHERE  Id = 6")
	statement.Exec()
	statement, _ = store.Db.Prepare("UPDATE Plant SET PlantDate = '2019-09-08'  WHERE Id = 12")
	statement.Exec()
	statement, _ = store.Db.Prepare("UPDATE Plant SET PlantDate = '2019-09-10'  WHERE Id = 24")
	statement.Exec()
	statement, _ = store.Db.Prepare("UPDATE Plant SET PlantDate = '2019-09-30'  WHERE Id = 30")
	statement.Exec()

	//Make  Empty Plant Pot
	statement, _ = store.Db.Prepare("UPDATE Plant SET Harvested = 1, PlantPosition = 0  WHERE Id = 24")
	statement.Exec()
	statement, _ = store.Db.Prepare("UPDATE Plant SET Harvested = 1, PlantPosition = 0  WHERE Id = 11")
	statement.Exec()
	statement, _ = store.Db.Prepare("UPDATE Plant SET Harvested = 1, PlantPosition = 0  WHERE Id = 12")
	statement.Exec()

	rows, _ = store.Db.Query("SELECT Id, Position, PlantType, AvailableSpots, TotalSpots from Module")
	var Position int
	var Id int
	var PlantType string
	var AvailableSpots int
	var TotalSpots int
	for rows.Next() {
		rows.Scan(&Id, &Position, &PlantType, &AvailableSpots, &TotalSpots)
		fmt.Println(strconv.Itoa(Id), strconv.Itoa(Position)+PlantType+strconv.Itoa(AvailableSpots)+strconv.Itoa(TotalSpots))
	}

	//	Create Sensor Table
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Create Sensor Table")
	statement, _ = store.Db.Prepare("DROP TABLE IF EXISTS SensorMeasurements")
	statement.Exec()
	statement, _ = store.Db.Prepare("CREATE TABLE IF NOT EXISTS SensorMeasurements (Id INTEGER PRIMARY KEY, Datetime TEXT, Temp REAL, LightIntensity REAL, Humidity REAL, WaterLevel REAL, WaterTemp REAL, WaterpH REAL)")
	statement.Exec()
	statement, _ = store.Db.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")

	yesterday := time.Date(2019, 11, 17, 20, 0, 0, 0, time.UTC)
	//light := 100
	temp := 0
	fmt.Println("then:", yesterday.Format(time.RFC3339))

	for i := 0; i < 24; i++ {
		light := float64(rand.Intn(100-70) + 70)
		light += 0.05
		statement.Exec(yesterday.Format(time.RFC3339), rand.Intn(20-15)+15, light, rand.Intn(30)+20, rand.Intn(35)+5, rand.Intn(20)+10, rand.Intn(8)+2)
		yesterday = yesterday.Add(time.Hour)
		light += 1
		temp += 1
	}

	// Create Time table
	statement, _ = store.Db.Prepare("DROP TABLE IF EXISTS TimeTable")
	statement.Exec()
	statement, _ = store.Db.Prepare("CREATE TABLE IF NOT EXISTS TimeTable (Id INTEGER PRIMARY KEY, TimeName TEXT, OnTime Text, OffTime Text, CurrentState INTEGER)")
	statement.Exec()
	statement, _ = store.Db.Prepare("INSERT INTO TimeTable (TimeName, OnTime, OffTime, CurrentState) VALUES (?, ?, ?, ?)")
	statement.Exec("Light", "16:54", "17:07", 1)
	statement.Exec("Pump", "16:42", "00:00", 5)

	return
}

func (store *Database) CreatePlantTypes() (err error) {
	statement, _ := store.Db.Prepare("INSERT INTO PlantType (Name, Growthtime) VALUES (?, ?)")

	statement.Exec("Lettuce", 42)
	statement.Exec("Basil", 21)

	return
}

func (store *Database) CreateModule(moduleNumber int, plantType string, availableSpots int) (err error) {
	statement, _ := store.Db.Prepare("INSERT OR IGNORE INTO Module (Position, PlantType, AvailableSpots, TotalSpots) VALUES (?, ? ,?, ?)")
	statement.Exec(moduleNumber, plantType, availableSpots, 6)

	return
}

func (store *Database) AddPlantToModule(moduleNumber int, plantPosition int, plantDate string) (err error) {
	statement, _ := store.Db.Prepare("INSERT OR IGNORE INTO Plant (Module, PlantPosition, PlantDate, Harvested) VALUES (?, ?, ?, ?)")
	statement.Exec(moduleNumber, plantPosition, plantDate, 0)
	return
}

func (store *Database) FillModuleWithPlants(moduleNumber int, plantAmount int) (err error) {
	// For each plant go 7 days back -> -7, -14...
	for i := 0; i < plantAmount; i++ {
		days := time.Duration(7 * (i + 1))
		daysBack := time.Hour * 24 * -days
		timeNow := time.Now()
		timePast := timeNow.Add(daysBack)
		timeStampString := timePast.Format("2006-01-02")
		store.AddPlantToModule(moduleNumber, i+1, timeStampString)
	}
	return
}

func (store *Database) CreateModuleWithPlants(moduleNumber int, plantType string) (err error) {
	store.CreateModule(moduleNumber, plantType, 0)
	store.FillModuleWithPlants(moduleNumber, 6)

	return
}

func (store *Database) CreateModuleWithFreeSpot(moduleNumber int, plantType string, freeSpots int) (err error) {
	store.CreateModule(moduleNumber, plantType, freeSpots)
	store.FillModuleWithPlants(moduleNumber, 6-freeSpots)

	return
}
