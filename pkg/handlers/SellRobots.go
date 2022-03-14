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
				writer.WriteHeader(http.StatusNoContent)
			} else {
				var updateRobot models.Robot
				if err3 := json.Unmarshal(body, &updateRobot); err3 != nil {
					log.Fatalln(err3)
					writer.WriteHeader(http.StatusBadRequest)
				} else {
					if robot.Count-updateRobot.Count < 0 {
						writer.WriteHeader(http.StatusBadRequest)
						writer.Header().Add("Content-Type", "application/json")
						json.NewEncoder(writer).Encode("There are fewer of them than you want")
					} else {
						var sale models.Sale

						robot.Count -= updateRobot.Count
						sale.CountRobots = updateRobot.Count
						sale.SellTime = time.Now()
						sale.RobotId = robot.Id
						sale.Profit = (robot.SellingPrice - robot.ManufacturingCost - robot.Allowance) * updateRobot.Count

						if save1 := h.DB.Save(&robot); save1.Error != nil {
							log.Fatalln(save1)
							writer.WriteHeader(http.StatusNotModified)
						} else if save2 := h.DB.Save(&sale); save2.Error != nil {
							log.Fatalln(save2)
							writer.WriteHeader(http.StatusNotModified)
						} else {
							writer.WriteHeader(http.StatusOK)
							writer.Header().Add("Content-Type", "application/json")
							json.NewEncoder(writer).Encode("Sold")
						}
					}
				}
			}
		}
	}
}
