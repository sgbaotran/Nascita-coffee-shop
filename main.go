package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "REST-API ", log.LstdFlags)

	mux := http.NewServeMux()

	mux.Handle("/", handlers.NewProduct(l))

	// mux.Handle("/goodbye", handlers.NewGoodbye(l))

	// mux.Handle("/products", handlers.NewProduct(l))

	server := &http.Server{
		Addr:         ":3030",
		Handler:      mux,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	l.Println("Somebody turned off", sig)

}
