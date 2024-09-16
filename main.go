package main

import (
	"log"
	"net/http"

	"github.com/Nico3012/webserver"
)

func main() {
	err := webserver.CreateWebServerAndCertificate(":443", "trusted/ca.pem", "trusted/key.pem", []string{"localhost"}, http.FileServer(http.Dir("html")))

	log.Fatalln(err)
}
