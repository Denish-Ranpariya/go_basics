package main

import (
	"log"
	"net/http"

	"github.com/Denish-Ranpariya/mongoapi/router"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
