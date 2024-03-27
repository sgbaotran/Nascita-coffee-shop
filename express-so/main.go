package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "REST-API ", log.LstdFlags)

	// mux := http.NewServeMux()

	// mux.Handle("/", handlers.NewProduct(l))

	ph := handlers.NewProduct(l)

	r := mux.NewRouter()

	getRouter := r.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := r.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(handlers.ValidateProductMiddleWare)

	postRouter := r.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(handlers.ValidateProductMiddleWare)

	server := &http.Server{
		Addr:         ":3030",
		Handler:      r,
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
