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
)

func init() {
	flag.Parse()
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	firmataAdaptor := firmata.NewAdaptor(*port)
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			val, err := sensor.Read()
			if err != nil {
				log.Println("Failed to read", err)
				return
			}
			cel := (5.0 * float64(val) * 100.0) / 1024
			log.Printf("Raw-value:%v Celsius:%.2f", val, cel)
		})
	}

	robot := gobot.NewRobot("tempSensorBot",
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
