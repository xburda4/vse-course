package main

import (
	"fmt"
	"log"
	"net/http"

	netchi "vse-course/transport/chi"
)

func main() {
	r := netchi.Initialize()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r); err != nil {
		log.Fatal(err)
	}
}
