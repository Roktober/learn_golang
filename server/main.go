package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"server/controller"
	"server/service"
	"server/text/processor"
	"sorted_map_task/ordered/vanil"
	"time"
)

func main() {
	waitForStop := time.Second * 5
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	re, err := regexp.Compile(`[^\w ]`)
	if err != nil {
		panic(err)
	}
	orderedMap := vanil.NewOrderedMap(1000)
	ignoreToken := make(map[string]int)
	textProcessor := processor.TextProcessor{OrderedMap: orderedMap, IgnoreToken: ignoreToken, TokenRe: re}
	app := controller.TopWordController{StopCh: c, Service: service.Service{TextProcessor: textProcessor}}

	r := mux.NewRouter()
	r.HandleFunc("/stat/{count}", app.Stat).Methods("GET")
	r.HandleFunc("/text", app.Text).Methods("POST")
	r.HandleFunc("/stop", app.Stop)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), waitForStop)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Error while stoping app")
	}
	os.Exit(0)
}
