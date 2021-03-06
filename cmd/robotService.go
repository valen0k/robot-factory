package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robot-factory/pkg/db"
	"robot-factory/pkg/handlers"
	"robot-factory/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	DB, err := db.Init(config)
	if err != nil {
		log.Fatalln(err)
	}

	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/robots", h.GetAllRobots).Methods(http.MethodGet)
	router.HandleFunc("/robots", h.AddRobot).Methods(http.MethodPost)
	router.HandleFunc("/robots/{id}", h.GetRobot).Methods(http.MethodGet)
	router.HandleFunc("/robots/{id}", h.UpdateRobot).Methods(http.MethodPut)
	router.HandleFunc("/robots/{id}", h.DeleteRobot).Methods(http.MethodDelete)
	router.HandleFunc("/robots/{id}/sell_robots", h.SellRobots).Methods(http.MethodPut)
	router.HandleFunc("/profit", h.GetProfit).Methods(http.MethodGet)
	//router.HandleFunc("/future_profit", h.GetFutureProfit).Methods(http.MethodGet)

	log.Println("API is Running")
	err = http.ListenAndServe(config.Port, router)
	if err != nil {
		log.Fatalln(err)
	}
}
