package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robot-factory/pkg/models"
	"strconv"
	"time"
)

func (h handler) GetProfit(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, err1 := strconv.Atoi(vars["year"])
	month, err2 := strconv.Atoi(vars["month"])
	if err1 != nil || err2 != nil {
		log.Fatalln(err1, err2)
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		var sales []models.Sale
		if find := h.DB.Find(&sales,
			"sell_time BETWEEN ? AND ?",
			strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-14 00:00:00.000000 +00:00",
			time.Now()); find.Error != nil {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			var profit int
			for i := 0; i < len(sales); i++ {
				profit += sales[i].Profit
			}
			writer.Header().Add("Content-Type", "application/json")
			json.NewEncoder(writer).Encode("The profit is " + strconv.Itoa(profit))
		}
	}
}
