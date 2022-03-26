package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"robot-factory/pkg/models"
	"strconv"
	"time"
)

func (h handler) GetProfit(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err1 := ioutil.ReadAll(request.Body)
	if err1 != nil {
		log.Println(err1)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var profitBody models.ProfitRequest
	if err2 := json.Unmarshal(body, &profitBody); err2 != nil {
		log.Println(err2)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	startTime := time.Date(profitBody.Year,
		time.Month(profitBody.Month),
		profitBody.Day,
		0, 0, 0, 0, time.Local)
	var finishTime time.Time
	if profitBody.AmountDays < 1 {
		finishTime = time.Now()
	} else {
		finishTime = startTime.Add(time.Hour*24*time.Duration(profitBody.AmountDays) - time.Second)
	}
	var profit1 int
	h.DB.Table("sales").Select("SUM((sell_price - cost) * count_robots)").
		Where("sell_time BETWEEN ? AND ?",
			startTime, finishTime).Row().Scan(&profit1)
	var profit2 int
	h.DB.Table("robots_warehouses").
		Select("SUM(warehouse_storage_cost)").
		Where("sale_id > 0 AND last_update_storage_cost BETWEEN ? AND ?",
			startTime, finishTime).Row().Scan(&profit2)
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("The profit is " + strconv.Itoa(profit1-profit2))
}
