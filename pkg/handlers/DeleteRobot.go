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
	id, err1 := strconv.Atoi(vars["id"])
	if err1 != nil {
		log.Println(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	// Delete that robot
	if tx := h.DB.Delete(models.Robot{}, id); tx.Error != nil {
		fmt.Println(tx.Error)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"status": "Deleted"})
}
