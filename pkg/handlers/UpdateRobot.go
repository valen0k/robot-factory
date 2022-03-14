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
		log.Fatalln(err1)
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		// Read request body
		defer request.Body.Close()
		body, err2 := ioutil.ReadAll(request.Body)
		if err2 != nil {
			log.Fatalln(err2)
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			// Find the robot by Id
			var robot models.Robot
			if first := h.DB.First(&robot, id); first.Error != nil {
				fmt.Println(first.Error)
				writer.WriteHeader(http.StatusNotFound)
			} else {
				var updateRobot models.Robot
				err3 := json.Unmarshal(body, &updateRobot)
				if err3 != nil {
					log.Fatalln(err2)
					writer.WriteHeader(http.StatusBadRequest)
				} else {
					robot.Type = updateRobot.Type
					robot.ManufacturingCost = updateRobot.ManufacturingCost
					robot.StorageCost = updateRobot.StorageCost
					robot.SellingPrice = updateRobot.SellingPrice
					robot.ManufacturingRate = updateRobot.ManufacturingRate

					if save := h.DB.Save(&robot); save.Error != nil {
						log.Fatalln(err2)
						writer.WriteHeader(http.StatusNotModified)
					} else {
						writer.WriteHeader(http.StatusOK)
						writer.Header().Add("Content-Type", "application/json")
						json.NewEncoder(writer).Encode("Updated")
					}
				}
			}
		}
	}
}
