package main

import (
	"flag"
)

var (
	flagParkappHost string
	flagParkappPort string
)

func init() {
	flag.StringVar(&flagParkappHost, "parkapp-host", "", "host to run parkapp HTTP server on")
	flag.StringVar(&flagParkappPort, "parkapp-port", "", "port to run parkapp HTTP server on")
}
