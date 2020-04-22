package db

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) GetPlantsPerType(p string) ([]*PlantsPerPlantType, error) {
	rets := m.Called(p)
	return rets.Get(0).([]*PlantsPerPlantType), rets.Error(1)
}

func (m *MockStore) GetHarvestablePlant(plantType *PlantType) (*PositionOnFarm, error) {
	rets := m.Called(plantType)
	return rets.Get(0).(*PositionOnFarm), rets.Error(1)
}

func (m *MockStore) HarvestDone(plantPosition *PositionOnFarm) (*Status, error) {
	rets := m.Called(plantPosition)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) Plant(plantType *PlantType) (int, error) {
	rets := m.Called(plantType)
	return rets.Get(0).(int), rets.Error(1)
}

func (m *MockStore) FinishPlanting(plantedModule *PlantedModule) (*Status, error) {
	rets := m.Called(plantedModule)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) DbSetup() error {
	rets := m.Called()
	return rets.Error(0)
}

func (m *MockStore) GetLastSensorEntry() (sensorData *SensorData, err error) {
	rets := m.Called()
	return rets.Get(0).(*SensorData), rets.Error(1)
}
func (m *MockStore) GetCatTreeData(module int) (plantInfo []*PlantInfoPerModule, err error) { //TODO Finish this
	return
}

func (m *MockStore) GetLightTimes() (currentTime *CurrentTime, err error) {
	rets := m.Called()
	return rets.Get(0).(*CurrentTime), rets.Error(1)
}

func (m *MockStore) InsertLightTimes(times *Times) (*Status, error) {
	rets := m.Called(times)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) GetPlantTypes() (plantTypes []*PlantTypes, err error) {
	rets := m.Called()
	return rets.Get(0).([]*PlantTypes), rets.Error(1)
}

func (m *MockStore) InsertModuleChanges(plantTypes *PlantTypes) (*Status, error) {
	rets := m.Called(plantTypes)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) GetKnownPlantTypes() (knownTypes []*KnownType, err error) {
	rets := m.Called()
	return rets.Get(0).([]*KnownType), rets.Error(1)
}

func (m *MockStore) InsertPumpTimes(pumpData *PumpData) (*Status, error) {
	rets := m.Called(pumpData)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) GetPumpTime() (pumpData *PumpData, err error) {
	rets := m.Called()
	return rets.Get(0).(*PumpData), rets.Error(1)
}

func (m *MockStore) AllPlantable() (plantableModules []*PlantableModules, err error) {
	rets := m.Called()
	return rets.Get(0).([]*PlantableModules), rets.Error(1)
}

func (m *MockStore) MassPlanting(plantedModules *PlantedModules) (status *Status, err error) {
	rets := m.Called(plantedModules)
	return rets.Get(0).(*Status), rets.Error(1)
}

func (m *MockStore) GetAllHarvestablePlant() (positionsOnFarm []*PositionOnFarm2, err error) {
	rets := m.Called()
	return rets.Get(0).([]*PositionOnFarm2), rets.Error(1)
}

func (m *MockStore) MassHarvest(plantPositions []PositionOnFarm) ( *Status,  error){
	rets := m.Called(plantPositions)
	return rets.Get(0).(*Status), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	FunctionStore = s
	return s
}
