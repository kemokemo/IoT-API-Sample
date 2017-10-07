package main

import (
	"flag"
	"log"
	"os"
	"sync"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

type sensorData struct {
	SensorList []sensor `json:"sensor_list"`
}

type sensor struct {
	Number       int     `json:"number"`
	Name         string  `json:"name" binding:"required"`
	TemperatureC float64 `json:"temp_c"`
}

var (
	addr     = flag.String("addr", ":5000", "the address of the API service")
	port     = flag.String("port", "/dev/ttyACM0", "the port of the Arduino")
	tempData sensorData
	tempLock sync.RWMutex
)

func init() {
	flag.Parse()
	tempData = sensorData{
		SensorList: []sensor{sensor{Number: 1, Name: "Kitchen"}},
	}
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	go runAPI(*addr)
	err := runRobot(*port)
	if err != nil {
		log.Println("Failed to start a robot", err)
		return exitCodeFailed
	}
	return exitCodeOK
}
