package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"robot-factory/pkg/models"
)

func (h handler) GetAllRobots(writer http.ResponseWriter, request *http.Request) {
	var robots []models.Robot
	if find := h.DB.Order("id asc").Find(&robots); find.Error != nil {
		log.Println(find.Error)
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(robots)
}
