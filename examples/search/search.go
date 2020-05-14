package main

import (
	"context"
	"crypto/tls"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	qradar "github.com/ilyaglow/go-qradar"
)

var (
	qradarLocation = flag.String("u", "", "QRadar Location URL")
	seckey         = flag.String("k", "", "SEC key to access the QRadar API")
	query          = flag.String("q", "", "Ariel SQL query")
	fieldsCSV      = flag.String("f", "-", "Fields to get (comma separated) as CSV data")
	window         = flag.Int("w", 500, "Number of events to search in batch")
	insecure       = flag.Bool("insecure", false, "Allow insecure TLS configuration")
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
		qradar.SetHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	qradar.SearchResultsWindow = *window
	scroller, _, err := qr.Ariel.ScrollByQuery(
		context.Background(),
		*query,
	)
	if err != nil {
		log.Fatal(err)
	}

	fields := []string{}
	if *fieldsCSV != "-" {
		fields = strings.Split(*fieldsCSV, ",")
	}

	if len(fields) == 0 {
		for scroller.Next(context.Background()) {
			fmt.Println(scroller.Result())
		}
		return
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll([][]string{fields})

	for scroller.Next(context.Background()) {
		result := scroller.Result()
		var selected []string
		for _, f := range fields {
			v, ok := result[f]
			if !ok {
				log.Fatalf("no value for %s in results %v", f, result)
			}

			selected = append(selected, strings.TrimSuffix(fmt.Sprintf("%v", v), "\n"))
		}
		w.WriteAll([][]string{selected})
	}
}
