package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	exitCodeOK int = iota
	exitCodeFailed
)

type sensorData struct {
	SensorList []sensor `json:"sensor-list"`
}

type sensor struct {
	Number       int     `json:"number"`
	Name         string  `json:"name" binding:"required"`
	TemperatureC float64 `json:"temp_c"`
}

var tempData sensorData

func init() {
	tempData = sensorData{
		SensorList: []sensor{
			sensor{Number: 1, Name: "Kitchen", TemperatureC: 25.86},
		},
	}
}

func main() {
	os.Exit(run(os.Args))
}

func run(args []string) int {
	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/sensors", sensorsGetEndpoint)
		v1.GET("/sensors/:sid", sensorIDGetEndpoint)
		v1.PUT("/sensors/:sid", sensorIDPutEndpoint)
	}

	err := r.Run(":5000")
	if err != nil {
		log.Println("Failed to start", err)
		return exitCodeFailed
	}
	return exitCodeOK
}

func sensorsGetEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, tempData)
}

func sensorIDGetEndpoint(c *gin.Context) {
	sid := c.Param("sid")
	id, err := parseSensorID(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tempData.SensorList[id-1])
}

func sensorIDPutEndpoint(c *gin.Context) {
	sid := c.Param("sid")
	id, err := parseSensorID(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data := sensor{}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"message": err.Error()})
		return
	}
	tempData.SensorList[id-1].Name = data.Name
	c.JSON(http.StatusOK, nil)
}

func parseSensorID(sid string) (int, error) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		return id, err
	}
	if id > len(tempData.SensorList) {
		return id, fmt.Errorf("id '%v' does not exist", id)
	}
	return id, nil
}
