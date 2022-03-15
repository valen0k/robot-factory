package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"robot-factory/pkg/models"
)

func (h handler) AddRobot(writer http.ResponseWriter, request *http.Request) {
	//	Read to request body
	defer request.Body.Close()
	body, err1 := ioutil.ReadAll(request.Body)
	if err1 != nil {
		log.Fatalln(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var robot models.Robot
	if err2 := json.Unmarshal(body, &robot); err2 != nil {
		log.Fatalln(err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	robot.Count = 0

	//	Append to the Robot mocks
	if create := h.DB.Create(&robot); create.Error != nil {
		log.Fatalln(create.Error)
		writer.WriteHeader(http.StatusNotModified)
		return
	}
	//	Send a 201 created response
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Created")
}
