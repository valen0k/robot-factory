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

	var profit int
	err3 := h.DB.Table("transaction_histories").Select("SUM((amount - manufacturing_cost) * count_robots)").
		Where("transaction = ? AND time BETWEEN ? AND ?",
			models.SALE, startTime, finishTime).Row().Scan(&profit)
	if err3 != nil {
		log.Println(err3)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var expenses int
	err4 := h.DB.Table("transaction_histories").Select("SUM((amount - manufacturing_cost) * count_robots)").
		Where("transaction = ? AND time BETWEEN ? AND ?",
			models.STORAGE, startTime, finishTime).Row().Scan(&expenses)
	if err4 != nil {
		log.Println(err4)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	res := map[string]string{
		"status": "success",
		"profit": strconv.Itoa(profit - expenses),
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(res)
}
