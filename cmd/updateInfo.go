package main

import (
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"robot-factory/pkg/db"
	"robot-factory/pkg/handlers"
	"syscall"
	"time"
)

const spec = "0 50 18 * * 1-5"

func main() {
	log.Println("Update info is Running")
	locationTime, err1 := time.LoadLocation("Europe/Moscow")
	if err1 != nil {
		log.Fatalln(err1.Error())
		return
	}
	DB, err2 := db.Init()
	if err2 != nil {
		log.Fatalln(err2.Error())
		return
	}
	h := handlers.New(DB)

	scheduler := cron.NewWithLocation(locationTime)

	defer scheduler.Stop()
	scheduler.AddFunc(spec, h.Trigger)
	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Update info is Stopping")
}
