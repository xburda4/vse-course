package main

import (
	"fmt"
	"log"
	"net/http"

	netchi "vse-course/transport/chi"
)

func main() {

	h := netchi.Initialize(8080)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", h.Port), h.Mux); err != nil {
		log.Fatal(err)
	}
}
