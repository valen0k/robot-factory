package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robot-factory/pkg/models"
	"strconv"
)

func (h handler) GetRobot(writer http.ResponseWriter, request *http.Request) {
	//	Read dynamic id parameter
	vars := mux.Vars(request)
	id, err1 := strconv.Atoi(vars["id"])
	if err1 != nil {
		log.Println(err1)
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
	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(robot)
}
