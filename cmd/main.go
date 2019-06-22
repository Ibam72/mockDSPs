package main

import (
	"log"
	"net/http"

	"github.com/Ibam72/mockdsps"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/mockDSPs/:dsp_id", mockdsps.POSTBidding)
	log.Fatal(http.ListenAndServe(":8080", router))
}
