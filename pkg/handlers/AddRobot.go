package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"robot-factory/pkg/models"
)

func (h handler) AddRobot(writer http.ResponseWriter, request *http.Request) {
	//	Read to request body
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var robot models.Robot
	json.Unmarshal(body, &robot)

	//	Append to the Robot mocks
	if create := h.DB.Create(&robot); create.Error != nil {
		fmt.Println(create.Error)
	}

	//	Send a 201 created response
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Created")

}
