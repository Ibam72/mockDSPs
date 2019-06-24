package main

import (
	"log"
	"net/http"
	"fmt"
	"os"

	"github.com/Ibam72/mockdsps"
	"github.com/julienschmidt/httprouter"
	flags "github.com/jessevdk/go-flags"
)

// Options is commandline options
type Options struct {
	Revision  bool     `short:"r" long:"revision" description:"Show revision information"`
	Port      string   `short:"p" long:"port" description:"port" default:"11115"`
}

func (o Options) port() string {
	return ":" + o.Port
}

var (
	opts     Options
	revision = ""
)

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	if opts.Revision {
		fmt.Printf("Revision:%v\n", revision)
		os.Exit(0)
	}
	router := httprouter.New()
	router.POST("/mockDSPs/:dsp_id", mockdsps.POSTBidding)
	log.Fatal(http.ListenAndServe(opts.port(), router))
}
