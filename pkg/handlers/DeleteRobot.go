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

func (h handler) DeleteRobot(writer http.ResponseWriter, request *http.Request) {
	//	Read dynamic id parameter
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalln(err)
	}

	// Find the robot by Id
	var robot models.Robot
	if first := h.DB.First(&robot, id); first.Error != nil {
		fmt.Println(first.Error)
		writer.WriteHeader(http.StatusNoContent)
	} else {
		// Delete that robot
		h.DB.Delete(&robot)

		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("Deleted")
	}

}
