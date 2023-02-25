package main

import (
	"fmt"
	"log"
	"net/http"

	nethttp "vse-course/transport/vanilla"
)

func main() {
	r := nethttp.Router()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r); err != nil {
		log.Fatal(err)
	}
}
