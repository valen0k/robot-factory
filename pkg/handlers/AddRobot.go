package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"robot-factory/pkg/models"
	"time"
)

func (h handler) AddRobot(writer http.ResponseWriter, request *http.Request) {
	//	Read to request body
	defer request.Body.Close()
	body, err1 := ioutil.ReadAll(request.Body)
	if err1 != nil {
		log.Println(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var robot models.Robot
	if err2 := json.Unmarshal(body, &robot); err2 != nil {
		log.Println(err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(robot.Type) < 3 ||
		robot.CountOfRobots < 0 ||
		robot.ManufacturingCost < 0 ||
		robot.ManufacturingRate < 0 ||
		robot.SellingPrice < 0 ||
		robot.StorageCost < 0 {
		log.Println("Bad Request")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	robot.LastUpdateNumberRobots = time.Now().Add(time.Hour * -24) // the time of the last update (minus a day for regular tasks)
	robot.LastUpdateStorageCost = time.Now().Add(time.Hour * -24)  // the time of the last update (minus a day for regular tasks)

	//	Append to the Robot mocks
	if create := h.DB.Create(&robot); create.Error != nil {
		log.Println(create.Error)
		writer.WriteHeader(http.StatusNotModified)
		return
	}
	//	Send a 201 created response
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"status": "Created"})
}
