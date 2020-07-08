package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MarcelCode/ROWA/src/db"
	"github.com/stretchr/testify/assert"
)

func TestPlantHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("Plant", &db.PlantType{}).Return(1, nil).Once()

	c, rec := InitialiseTestServer(http.MethodPost, "/plant/get-position")

	expected := 1

	if assert.NoError(t, PlantHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody int
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}

func TestFinishPlantingHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("FinishPlanting", &db.PlantedModule{}).Return(
		&db.Status{Message: "Planting Done"}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodPost, "/plant/finish")

	expected := &db.Status{Message: "Planting Done"}

	if assert.NoError(t, FinishPlantingHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.Status
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}


func TestMassPlantingHandler(t *testing.T) {
	mockStore := db.InitMockStore()
	mockStore.On("MassPlanting", &db.PlantedModules{}).Return(
		&db.Status{Message: "Planting Done"}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodPost, "/plant/plant-all")

	expected := &db.Status{Message: "Planting Done"}

	if assert.NoError(t, MassPlantingHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.Status
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}