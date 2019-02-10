package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	address = flag.String("address", ":9000", "address to exporter metrics")
)

func main() {
	flag.Parse()

	http.Handle("/", promhttp.Handler())

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Println("listening on", *address)

	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}
