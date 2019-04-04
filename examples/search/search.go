package main

import (
	"context"
	"flag"
	"log"
	"os"

	qradar "github.com/ilyaglow/go-qradar"
)

var (
	qradarLocation = flag.String("u", "", "QRadar Location URL")
	seckey         = flag.String("k", "", "SEC key to access the QRadar API")
	query          = flag.String("q", "", "Ariel SQL query")
)

func main() {
	flag.Parse()
	if *qradarLocation == "" || *seckey == "" || *query == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	qr, err := qradar.NewClient(
		*qradarLocation,
		qradar.SetSECKey(*seckey),
	)
	if err != nil {
		log.Fatal(err)
	}

	scroller, _, err := qr.Ariel.SearchByQuery(
		context.Background(),
		*query,
	)
	if err != nil {
		log.Fatal(err)
	}

	for scroller.Next(context.Background()) {
		log.Println(scroller.Result())
	}
}
