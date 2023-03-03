package main

import (
	"fmt"
	"log"
	"net/http"

	nethttp "vse-course/transport/vanilla"
)

func main() {

	h := nethttp.Initialize(8080)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.Mux); err != nil {
		log.Fatal(err)
	}
}
