package main

import (
	"apod/database"
	"apod/service"
	"apod/worker"
	"context"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	database.Initialize()

	ctx := context.Background()
	w := worker.NewWorker(ctx)
	w.Start()

	log.Println("HERE")
	apodService := &service.ApodService{}
	//http.HandleFunc("/apod/save", apodService.SaveApod)
	http.HandleFunc("/apod", apodService.GetApodByDate)
	http.HandleFunc("/apods", apodService.GetAllApods)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
