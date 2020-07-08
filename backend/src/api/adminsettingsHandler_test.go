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
	mockStore.On("GetLightTimes").Return(
		&db.CurrentTime{TimeOn: "18:50", TimeOff: "22:00", State: 1}, nil).Once()

	c, rec := InitialiseTestServer(http.MethodGet, "/adminSettings/get-light")

	expected := &db.CurrentTime{TimeOn: "18:50", TimeOff: "22:00", State: 1}

	if assert.NoError(t, GetLightTimes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		requestBody := &db.CurrentTime{}
		err := json.NewDecoder(rec.Body).Decode(&requestBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expected, requestBody)
	}

	mockStore.AssertExpectations(t)
}