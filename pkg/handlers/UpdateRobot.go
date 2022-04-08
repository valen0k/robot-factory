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
)

func (h handler) UpdateRobot(writer http.ResponseWriter, request *http.Request) {

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
	err3 := json.Unmarshal(body, &updateRobot)
	if err3 != nil {
		log.Println(err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if (len(updateRobot.Type) < 3 && len(updateRobot.Type) != 0) ||
		updateRobot.CountOfRobots < 0 ||
		updateRobot.ManufacturingCost < 0 ||
		updateRobot.ManufacturingRate < 0 ||
		updateRobot.SellingPrice < 0 ||
		updateRobot.StorageCost < 0 {
		log.Println("Bad Request")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(updateRobot.Type) != 0 {
		robot.Type = updateRobot.Type
	}
	if updateRobot.ManufacturingCost != 0 {
		robot.ManufacturingCost = updateRobot.ManufacturingCost
	}
	if updateRobot.StorageCost != 0 {
		robot.StorageCost = updateRobot.StorageCost
	}
	if updateRobot.SellingPrice != 0 {
		robot.SellingPrice = updateRobot.SellingPrice
	}
	if updateRobot.ManufacturingRate != 0 {
		robot.ManufacturingRate = updateRobot.ManufacturingRate
	}

	if save := h.DB.Save(&robot); save.Error != nil {
		log.Println(save.Error)
		writer.WriteHeader(http.StatusNotModified)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"status": "Updated"})
}
