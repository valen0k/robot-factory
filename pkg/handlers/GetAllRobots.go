package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"robot-factory/pkg/models"
)

func (h handler) GetAllRobots(writer http.ResponseWriter, request *http.Request) {
	var robots []models.Robot
	if find := h.DB.Find(&robots); find.Error != nil {
		fmt.Println(find.Error)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(robots)

}
