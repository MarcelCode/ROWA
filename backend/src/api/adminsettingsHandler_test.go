package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MarcelCode/ROWA/src/db"
	"github.com/stretchr/testify/assert"
	_ "golang.org/x/net/context"
)

func TestGetLightTimes(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("GetLightTimes", "harvestable").Return(
		[]*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/harvestable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}

	if assert.NoError(t, GetHarvestablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}

func TestGetPlantablePlantsHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("GetPlantsPerType", "plantable").Return(
		[]*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/plantable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}

	if assert.NoError(t, GetPlantablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}
