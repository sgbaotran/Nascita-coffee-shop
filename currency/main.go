package main

import (
	"net"
	"os"

	protos "github.com/sgbaotran/Nascita-coffee-shop/currency/protos/currency"
	"github.com/sgbaotran/Nascita-coffee-shop/currency/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	cs := server.NewServer(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Something goes wrong")
		os.Exit(1)
	}
	gs.Serve(l)

}
