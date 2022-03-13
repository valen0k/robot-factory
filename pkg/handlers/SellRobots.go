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

func (h handler) SellRobots(writer http.ResponseWriter, request *http.Request) {
	//	Read dynamic id parameter
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	// Read request body
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Find the robot by Id
	var robot models.Robot
	if first := h.DB.First(&robot, id); first.Error != nil {
		fmt.Println(first.Error)
		writer.WriteHeader(http.StatusNoContent)
	} else {
		var updateRobot models.Robot
		json.Unmarshal(body, &updateRobot)

		robot.Count -= updateRobot.Count

		h.DB.Save(&robot)

		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode("Updated")
	}
}
