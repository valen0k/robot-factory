package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"robot-factory/pkg/db"
	"robot-factory/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/robots", h.GetAllRobots).Methods(http.MethodGet)
	router.HandleFunc("/robots/{id}", h.GetRobot).Methods(http.MethodGet)
	router.HandleFunc("/robots", h.AddRobot).Methods(http.MethodPost)
	router.HandleFunc("/robots/{id}", h.UpdateRobot).Methods(http.MethodPut)
	router.HandleFunc("/robots/{id}", h.DeleteRobot).Methods(http.MethodDelete)

	log.Println("API is Running")
	http.ListenAndServe(":8081", router)
}
