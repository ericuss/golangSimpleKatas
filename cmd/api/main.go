package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	beersData "github.com/ericuss/golangSimpleKatas/cmd/beers-data"
	beer "github.com/ericuss/golangSimpleKatas/pkg"
	server "github.com/ericuss/golangSimpleKatas/pkg/server"
	mongo "github.com/ericuss/golangSimpleKatas/pkg/storage/mongo"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some beers")
	flag.Parse()

	var beers map[int]*beer.Beer
	if *withData {
		beers = beersData.Beers
	}

	repo := mongo.NewBeerRepository(beers)
	s := server.New(repo)

	fmt.Println("The beers server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
