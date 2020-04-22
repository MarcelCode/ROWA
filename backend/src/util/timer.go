package util

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/jasonlvhit/gocron"
)

type Times struct {
	TimeOn  string
	TimeOff string
}

var light = gocron.NewScheduler()

func Runner() {
	<-light.Start()
}

func LightTimesRenew() {

	light.Remove(sensor.LightSwitch)
	light.Remove(sensor.LightSwitch)

	sqlQuery := `SELECT OnTime, OffTime
				 FROM TimeTable
				 WHERE ID = 1`
	database, _ := sql.Open("sqlite3", "./rowa.db")

	rows, err := database.Query(sqlQuery)

	if err != nil {
		return
	}

	restartTime := new(Times)
	restartTime = &Times{}
	rows.Next()
	err = rows.Scan(&restartTime.TimeOn, &restartTime.TimeOff)

	TimeOnArray := strings.Split(restartTime.TimeOn, ":")
	TimeOffArray := strings.Split(restartTime.TimeOff, ":")

	TimeOnHour, _ := strconv.Atoi(TimeOnArray[0])
	TimeOnMinute, _ := strconv.Atoi(TimeOnArray[1])

	TimeOffHour, _ := strconv.Atoi(TimeOffArray[0])
	TimeOffMinute, _ := strconv.Atoi(TimeOffArray[1])
	fmt.Println(restartTime.TimeOff)
	fmt.Println(restartTime.TimeOn)
	tOn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOnHour, TimeOnMinute, 0, 0, time.Local)
	tOff := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOffHour, TimeOffMinute, 0, 0, time.Local)

	light.Every(1).Day().At(restartTime.TimeOn).From(&tOn).Do(sensor.LightSwitch, true)
	light.Every(1).Day().At(restartTime.TimeOff).From(&tOff).Do(sensor.LightSwitch, false)
	rows.Close()
}

func PumpTimesRenew() {
	light.Remove(sensor.TriggerPump)
	light.Remove(sensor.TriggerPump)

	sqlQuery := `SELECT OnTime, CurrentState
				 FROM TimeTable
				 WHERE ID = 2`
	database, _ := sql.Open("sqlite3", "./rowa.db")

	rows, err := database.Query(sqlQuery)

	if err != nil {
		return
	}

	restartTime := new(Times)
	restartTime = &Times{}
	rows.Next()
	var PumpTime int

	//Save On Time and Save Minutes and hours from it
	err = rows.Scan(&restartTime.TimeOn, &PumpTime)
	TimeOnArray := strings.Split(restartTime.TimeOn, ":")
	TimeOnHour, _ := strconv.Atoi(TimeOnArray[0])
	TimeOffHour, _ := strconv.Atoi(TimeOnArray[0])
	TimeOnMinute, _ := strconv.Atoi(TimeOnArray[1])

	TimeOffMinute, _ := (strconv.Atoi(TimeOnArray[1]))
	if PumpTime+TimeOffMinute >= 10 && PumpTime+TimeOffMinute < 60 {
		TimeOffMinute = TimeOffMinute + PumpTime
		restartTime.TimeOff = strconv.Itoa(TimeOffHour) + ":" + strconv.Itoa(TimeOffMinute)

	} else if PumpTime+TimeOffMinute < 10 {
		TimeOffMinute = TimeOffMinute + PumpTime
		restartTime.TimeOff = strconv.Itoa(TimeOffHour) + ":" + "0" + strconv.Itoa(TimeOffMinute)

	} else if PumpTime+TimeOffMinute < 70 {
		TimeOffMinute = TimeOffMinute + PumpTime - 60
		TimeOffHour = HourAdder(TimeOffHour)
		restartTime.TimeOff = strconv.Itoa(TimeOffHour) + ":" + "0" + strconv.Itoa(TimeOffMinute)

	} else {
		TimeOffMinute = TimeOffMinute + PumpTime - 60
		TimeOffHour = HourAdder(TimeOffHour)
		restartTime.TimeOff = strconv.Itoa(TimeOffHour) + ":" + strconv.Itoa(TimeOffMinute)
	}

	fmt.Println(restartTime.TimeOff)
	fmt.Println(restartTime.TimeOn)
	tOn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOnHour, TimeOnMinute, 0, 0, time.Local)
	tOff := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), TimeOffHour, TimeOffMinute, 0, 0, time.Local)

	light.Every(1).Day().At(restartTime.TimeOn).From(&tOn).Do(sensor.TriggerPump, true)
	light.Every(1).Day().At(restartTime.TimeOff).From(&tOff).Do(sensor.TriggerPump, false)
	rows.Close()
}

func HourAdder(TimeOffHour int) int {
	if TimeOffHour+1 == 24 {
		HoursPlusOne := 0
		return HoursPlusOne
	} else {
		HoursPlusOne := TimeOffHour + 1
		return HoursPlusOne
	}
}
