package main_test

import (
	"api"
	"database/sql"
	"db"
	"encoding/json"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func InitialiseTestServer(httpMethod string, url string) (c echo.Context, rec *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(httpMethod, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	return
}

func TestSensorDataStatus(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/sensor-data")

	if assert.NoError(t, api.GetSensorDataHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSensorData(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/sensor-data")

	expected := &db.SensorData{Datetime: "2019-11-18T19:00:00Z", Temp: 23, LightIntensity: 123}
	if assert.NoError(t, api.GetSensorDataHandler(c)) {
		requestBody := &db.SensorData{}
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetHarvestablePlantsStatus(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/harvestable-plants")

	if assert.NoError(t, api.GetHarvestablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetHarvestablePlantsData(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/harvestable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 2}, {"Lettuce", 2}}

	if assert.NoError(t, api.GetHarvestablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetPlantablePlantsData(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/plantable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 1}}

	if assert.NoError(t, api.GetPlantablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetPlantablePlantsFinal(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/plantable-plants")

	var expected []*db.PlantsPerPlantType

	if assert.NoError(t, api.GetPlantablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetHarvestablePlantsDataOneLess(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/dashboard/harvestable-plants")

	expected := []*db.PlantsPerPlantType{{"Basil", 1}, {"Lettuce", 2}}

	if assert.NoError(t, api.GetHarvestablePlantsHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody []*db.PlantsPerPlantType
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetHarvestablePlant(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/harvest/get-plant?plant_type=Basil")

	expected := &db.PositionOnFarm{6, 1}

	if assert.NoError(t, api.GetHarvestablePlantHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.PositionOnFarm
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestGetPlantablePlant(t *testing.T) {
	c, rec := InitialiseTestServer(http.MethodGet, "/plant/get-position?plant_type=Basil")

	expected := 1

	if assert.NoError(t, api.PlantHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody int
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestHarvestDone(t *testing.T) {
	JSON := `{"plant_position":6,"module_position":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/harvest/harvestdone", strings.NewReader(JSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.Status{Message: "harvest done"}

	if assert.NoError(t, api.HarvestDoneHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.Status
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestPlantingDone(t *testing.T) {
	JSON := `{"planted_module":1}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/plant/finish", strings.NewReader(JSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expected := &db.Status{Message: "Planting Done"}

	if assert.NoError(t, api.FinishPlantingHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var requestBody *db.Status
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}
}

func TestMain(m *testing.M) {
	database, err := sql.Open("sqlite3", "mock.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	db.InitStore(&db.Database{Db: database})
	db.FunctionStore.DbSetup()

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestDashboard(t *testing.T) {
	t.Run("status=sensor", TestSensorDataStatus)
	t.Run("data=sensor", TestSensorData)
	t.Run("status=harvestable_plants", TestGetHarvestablePlantsStatus)
	t.Run("data=harvestable_plants", TestGetHarvestablePlantsData)
}

func TestHarvestPlant(t *testing.T) {
	t.Run("harvestable_plants", TestGetHarvestablePlantsData)
	t.Run("harvestable_plant_position", TestGetHarvestablePlant)
	t.Run("harvest_done", TestHarvestDone)
	t.Run("check_harvestable_count", TestGetHarvestablePlantsDataOneLess)
	t.Run("check_plantable_count", TestGetPlantablePlantsData)
	t.Run("get_plant_module", TestGetPlantablePlant)
	t.Run("planting_done", TestPlantingDone)
	t.Run("check_plantable_count_final", TestGetPlantablePlantsFinal)
}
