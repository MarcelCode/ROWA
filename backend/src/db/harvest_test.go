package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func (s *StoreSuite) TestGetHarvestablePlant() {
	s.store.CreatePlantTypes()
	s.store.CreateModuleWithPlants(1, "Basil")
	s.store.CreateModuleWithPlants(2, "Lettuce")

	inputs := []*PlantType{{Name: "Basil"}, {Name: "Lettuce"}}
	expectedOutcome := []*PositionOnFarm{{3, 1}, {6, 2}}

	for i, element := range inputs {
		harvestablePlant, err := s.store.GetHarvestablePlant(element)
		if err != nil {
			s.T().Fatal(err)
		}

		assert.Equal(s.T(), harvestablePlant, expectedOutcome[i])
	}
}

func (s *StoreSuite) TestHarvestDone() {
	s.store.CreatePlantTypes()
	s.store.CreateModuleWithPlants(1, "Basil")
	s.store.CreateModuleWithPlants(2, "Lettuce")

	harvestDone, err := s.store.HarvestDone(&PositionOnFarm{6, 1})
	harvestDone2, _ := s.store.HarvestDone(&PositionOnFarm{6, 2})
	if err != nil {
		s.T().Fatal(err)
	}
	expected := &Status{"harvest done"}
	assert.Equal(s.T(), harvestDone, expected,harvestDone2)

	plantablePlants, err := s.store.GetPlantsPerType("plantable")
	if err != nil {
		s.T().Fatal(err)
	}

	dbChange := []*PlantsPerPlantType{{"Basil", 1}, {"Lettuce", 1}}
	assert.Equal(s.T(),dbChange, plantablePlants)
}
