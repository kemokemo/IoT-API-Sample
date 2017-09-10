package main

import (
	"flag"
	"log"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

var (
	port = flag.String("port", "/dev/ttyACM0", "the port of the Arduino")
	pin  = flag.String("pin", "0", "the pin number of the analog")
)

func init() {
	flag.Parse()
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	firmataAdaptor := firmata.NewAdaptor(*port)
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, *pin)

	work := func() {
		gobot.Every(1*time.Second, func() {
			val, err := sensor.Read()
			if err != nil {
				log.Println("Failed to read", err)
				return
			}
			log.Printf("Raw-value:%v", val)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Println("Failed to start a robot", err)
		return exitCodeFailed
	}
	return exitCodeOK
}
