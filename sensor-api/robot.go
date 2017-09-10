package main

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/platforms/firmata"
)

func runRobot(port string) error {
	firmataAdaptor := firmata.NewAdaptor(port)
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0")

	work := func() {
		gobot.Every(1*time.Second, func() {
			val, err := sensor.Read()
			if err != nil {
				log.Println("Failed to read", err)
				return
			}
			cel := (5.0 * float64(val) * 100.0) / 1024.0
			tempLock.Lock()
			tempData.SensorList[0].TemperatureC = cel
			tempLock.Unlock()
		})
	}

	robot := gobot.NewRobot("tempAPIBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor},
		work,
	)

	err := robot.Start()
	if err != nil {
		return err
	}
	return nil
}
