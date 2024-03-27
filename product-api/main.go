package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sgbaotran/Nascita-coffee-shop/product-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "REST-API ", log.LstdFlags)

	// mux := http.NewServeMux()

	// mux.Handle("/", handlers.NewProduct(l))

	ph := handlers.NewProduct(l)

	r := mux.NewRouter()

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(handlers.ValidateProductMiddleWare)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(handlers.ValidateProductMiddleWare)

	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct)
	deleteRouter.Use(handlers.ValidateProductMiddleWare)

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
