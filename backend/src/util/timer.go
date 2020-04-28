package util

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
	"os"
    log "github.com/sirupsen/logrus"
	"github.com/MarcelCode/ROWA/src/sensor"
	"github.com/jasonlvhit/gocron"
)

type Times struct {
	TimeOn  string
	TimeOff string
}

var light = gocron.NewScheduler()
var logs = gocron.NewScheduler()

func Runner() {
	<-light.Start()
}

func RunnerLog() {
	<-logs.Start()
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

func LogNameIntervall(){
	logs.Remove(LogRenew)
    logs.Every(1).Day().At("00:00").From(gocron.NextTick()).Do(LogRenew)
	//logs.Every(2).Seconds().Do(LogRenew)
	
}

func LogRenew() {
	var filename string = "logfile-backend-"+ time.Now().Format("2006-01-02") +".log"
    // Create the log file if doesn't exist. And append to it if it already exists.
    f, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
    Formatter := new(log.TextFormatter)
    // You can change the Timestamp format. But you have to use the same date and time.
    // "2006-02-02 15:04:06" Works. If you change any digit, it won't work
    // ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
    Formatter.TimestampFormat = "02-01-2006 15:04:05"
    Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
    if err != nil {
        // Cannot open log file. Logging to stderr
        fmt.Println(err)
    }else{
        log.SetOutput(f)
	}
	
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
