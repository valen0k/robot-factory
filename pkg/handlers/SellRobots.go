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
		log.Fatalln(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Read request body
	defer request.Body.Close()
	body, err2 := ioutil.ReadAll(request.Body)
	if err2 != nil {
		log.Fatalln(err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Find the robot by Id
	var robot models.Robot
	if first := h.DB.First(&robot, id); first.Error != nil {
		fmt.Println(first.Error)
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	var updateRobot models.Robot
	if err3 := json.Unmarshal(body, &updateRobot); err3 != nil {
		log.Fatalln(err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if robot.Count-updateRobot.Count < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("There are fewer of them than you want")
		return
	}
	var sale models.Sale

	robot.Count -= updateRobot.Count
	sale.RobotId = robot.Id
	sale.CountRobots = updateRobot.Count
	sale.ManufacturingCost = robot.ManufacturingCost
	sale.WarehouseStorageCost = robot.WarehouseStorageCost
	sale.SellingPrice = robot.SellingPrice
	sale.SellTime = time.Now()
	//sale.Profit = (robot.SellingPrice - robot.ManufacturingCost - robot.WarehouseStorageCost) * updateRobot.Count

	save1 := h.DB.Save(&robot)
	if save1.Error != nil {
		log.Fatalln(save1)
		writer.WriteHeader(http.StatusNotModified)
		return
	}
	save2 := h.DB.Save(&sale)
	if save2.Error != nil {
		log.Fatalln(save2)
		writer.WriteHeader(http.StatusNotModified)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Sold")
}
