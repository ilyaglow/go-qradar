package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	qradar "github.com/ilyaglow/go-qradar"
)

var (
	qradarLocation = flag.String("u", "", "QRadar Location URL")
	seckey         = flag.String("k", "", "SEC key to access the QRadar API")
	filter         = flag.String("filter", "", "Filter string")
	domain         = flag.String("domain", "", "Domain description to filter (could be a part of it)")
)

func main() {
	flag.Parse()
	if *qradarLocation == "" || *seckey == "" {
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

	if *domain != "" {
		domains, err := qr.Config.Domains(context.Background(), "", "", 0, 100)
		if err != nil {
			log.Fatal(err)
		}

		var id int
		for i := range domains {
			if strings.Contains(domains[i].Description, *domain) {
				id = domains[i].ID
				break
			}
		}

		if id != 0 {
			*filter = fmt.Sprintf("%s and domain_id=%d", *filter, id)
		}
	}

	offs, err := qr.SIEM.Offenses(context.Background(), "", *filter, "-start_time", 0, 100)
	if err != nil {
		log.Fatal(err)
	}

	for i := range offs {
		fmt.Printf("%s %s\n", offs[i].Description, time.Unix(int64(offs[i].StartTime)/1000, 0))
	}
}
