package main

import (
	"flag"
	"git.rabobank.nl/happy-endpoint/handlers"
	"git.rabobank.nl/happy-endpoint/helpers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Variables to identiy the build
var (
	CommitHash string
	VersionTag string
	BuildTime  string
)

func main() {
	var (
		httpAddr = flag.String("http", "0.0.0.0:7000", "HTTP server address")
	)
	flag.Parse()

	log.Println("Starting Happy Endpoint")
	log.Printf("The version is: %s; the commit hash is: %s. Build time is: %s", VersionTag, CommitHash, helpers.ParseBuildTime(BuildTime).Format(time.RFC1123))
	log.Printf("listening on address %s", *httpAddr)

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HelloHandler)
	//r.HandleFunc("/version", handlers.VersionHandler)

	log.Fatal(http.ListenAndServe(*httpAddr, r))
}
