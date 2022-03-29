package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"robot-factory/pkg/models"
	"strconv"
	"time"
)

func (h handler) SellRobots(writer http.ResponseWriter, request *http.Request) {
	//	Read dynamic id parameter
	vars := mux.Vars(request)
	id, err1 := strconv.Atoi(vars["id"])
	if err1 != nil {
		log.Println(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Read request body
	defer request.Body.Close()
	body, err2 := ioutil.ReadAll(request.Body)
	if err2 != nil {
		log.Println(err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Find the robot by Id
	var robot models.Robot
	if first := h.DB.First(&robot, id); first.Error != nil {
		fmt.Println(first.Error)
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	var updateRobot models.Robot
	if err3 := json.Unmarshal(body, &updateRobot); err3 != nil {
		log.Println(err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if robot.Count-updateRobot.Count < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("There are fewer of them than you want")
		return
	}

	var history models.TransactionHistory
	now := time.Now()

	robot.Count -= updateRobot.Count
	history.Transaction = models.SALE
	history.RobotId = robot.Id
	history.CountRobots = updateRobot.Count
	history.Amount = robot.SellingPrice
	history.ManufacturingCost = robot.ManufacturingCost
	history.Time = now

	save1 := h.DB.Save(&robot)
	if save1.Error != nil {
		log.Println(save1)
		writer.WriteHeader(http.StatusNotModified)
		return
	}
	save2 := h.DB.Save(&history)
	if save2.Error != nil {
		log.Println(save2)
		writer.WriteHeader(http.StatusNotModified)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Sold")
}
