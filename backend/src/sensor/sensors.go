package sensor

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	//"github.com/jacobsa/go-serial/serial"
	"github.com/tarm/serial"
)

/*Serial Port Configs
/dev/ttyACM0 - Rasp
/dev/cu.usbmodem1434301 Macbook
COM5 windows
*/

func SetupSerialConnection() (s *serial.Port, err error) {
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Print(err)
		return
	}

	return s, err

}

func ActivateModuleLight(moduleNumber int) {
	// Create String from Module Number and send to connection
	moduleString := strconv.Itoa(moduleNumber)

	// Give Connection time to send Data
	WriteToCh(moduleString)
}
func TriggerPump(state bool) {

	if state {
		WriteToCh("90")
		fmt.Println("Pump on")
	} else {
		WriteToCh("91")
		fmt.Println("Pump off")
	}

}

func TriggerAirStone(state bool) {
	if state {
		WriteToCh("70")
		fmt.Println("Airstone on")
	} else {
		WriteToCh("71")
		fmt.Println("Airstone off")
	}
}

func DeactivateModuleLight() {
	WriteToCh("99")
}

func LightSwitch(state bool) {
	fmt.Println("Light switch triggered")

	//send turn off or on to arduino
	if state {
		WriteToCh("80")
		fmt.Println("Light on")
		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 1 WHERE ID = 1")
		statement.Exec()
		database.Close()
	} else {
		WriteToCh("81")
		fmt.Println("Light off")
		//change DB light State
		database, _ := sql.Open("sqlite3", "./rowa.db")
		statement, _ := database.Prepare("UPDATE TimeTable SET CurrentState= 0 WHERE ID = 1")
		statement.Exec()
		database.Close()
	}

}
func ReadFakeSensorData() {
	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()
	for {
		datetime := time.Now()
		temp := rand.Float32()*1 + 22
		lightIntensity := rand.Float32()*1 + 35
		humidity := rand.Float32()*10 + 30
		waterLevel := rand.Float64()*4 + 13
		waterTemp := rand.Float32()*1 + 22
		waterpH := rand.Float32() + 6
		datetimeStr := datetime.UTC().Format(time.RFC3339)
		fmt.Println(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
		//Sending notification
		CheckReadings(datetime, waterLevel)
		statement.Exec(datetimeStr, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)

		time.Sleep(2 * time.Second)
	}

}
func ReadSensorData(s *serial.Port) {
	var serialString string

	database, _ := sql.Open("sqlite3", "./rowa.db")
	statement, _ := database.Prepare("INSERT OR IGNORE INTO SensorMeasurements (Datetime, Temp, LightIntensity, Humidity, WaterLevel, WaterTemp, WaterpH) VALUES (?, ?, ?, ?, ?, ?, ?)")
	defer database.Close()

	for {
		buf := make([]byte, 128)
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		serialIncome := string(buf[:n])
		serialString += serialIncome
		//log.Println("Serial string: ", serialString)
		if strings.HasSuffix(serialString, "\n") {
			datetime := time.Now()
			raw_string := strings.TrimSuffix(serialString, "\r\n")
			data_array := strings.Split(raw_string, ",")
			if len(data_array) >= 6 {
				temp, err1 := strconv.ParseFloat(data_array[0], 32)
				lightIntensity, err2 := strconv.ParseFloat(data_array[1], 32)
				humidity, err3 := strconv.ParseFloat(data_array[2], 32)
				waterLevel, err4 := strconv.ParseFloat(data_array[3], 32)
				waterTemp, err5 := strconv.ParseFloat(data_array[4], 32)
				waterpH, err6 := strconv.ParseFloat(data_array[5], 32)
				if err1 == nil && err2 == nil && err4 == nil && err3 == nil && err5 == nil && err6 == nil {
					//log.Println(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
					//Sending notification
					CheckReadings(datetime, waterLevel)
					datetime := datetime.UTC().Format(time.RFC3339)
					//Writing to local db
					statement.Exec(datetime, temp, lightIntensity, humidity, waterLevel, waterTemp, waterpH)
					
				}

			}
			serialString = ""
		}

	}

}
