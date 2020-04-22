package db

import "database/sql"

type Store interface {
	GetPlantsPerType(p string) ([]*PlantsPerPlantType, error)
	GetLastSensorEntry() (*SensorData, error)
	GetHarvestablePlant(*PlantType) (*PositionOnFarm, error)
	HarvestDone(*PositionOnFarm) (*Status, error)
	Plant(*PlantType) (int, error)
	FinishPlanting(*PlantedModule) (*Status, error)
	GetLightTimes() (*CurrentTime, error)
	InsertLightTimes(*Times) (*Status, error)

	InsertModuleChanges(*PlantTypes) (*Status, error)

	GetPlantTypes() ([]*PlantTypes, error)
	DbSetup() error
	GetCatTreeData(module int) ([]*PlantInfoPerModule, error)
	GetKnownPlantTypes() ([]*KnownType, error)

	GetPumpTime() (*PumpData, error)
	InsertPumpTimes(*PumpData) (*Status, error)

	AllPlantable() ([]*PlantableModules, error)
	MassPlanting(*PlantedModules) (*Status, error)

	GetAllHarvestablePlant() ([]*PositionOnFarm2, error)
	MassHarvest([]PositionOnFarm) (*Status, error)
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type Database struct {
	Db *sql.DB
}

// The FunctionStore variable is a global level variable that will be available for
// use throughout our application code
var FunctionStore Store

/*
We will need to call the InitStore method to initialize the FunctionStore.
*/
func InitStore(s Store) {
	FunctionStore = s
}
