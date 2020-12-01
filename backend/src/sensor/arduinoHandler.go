package sensor

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
)

var ch = make(chan string, 100)

func WriteToCh(input string) {
	ch <- input
}

func ArduinoLoop(s *serial.Port) {

	log.Println("Starting loop")
	_, err := s.Write([]byte("80"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("81"))
	time.Sleep(time.Second * 1)
	_, err = s.Write([]byte("80"))
	time.Sleep(time.Second * 1)

	for {

		for v := range ch {
			log.Println("Writing ", v, "to arduino")
			_, err = s.Write([]byte(v))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)

		}
	}
}
