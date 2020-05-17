package main

import (
	"fmt"
	"log"
	"os"

	"blackhatgo/ch-3/shodan/shodan"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: shodan <command>\n\nCommand options:\n- Blank (provide search query)\n- ip <ip-address>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

	if len(os.Args) == 2 {

		hostSearch, err := s.HostSearch(os.Args[1])
		if err != nil {
			log.Panicln(err)
		}

		for _, host := range hostSearch.Matches {
			fmt.Printf("%18s%8d\n", host.IPString, host.Port)
		}
	} else if os.Args[1] == "ip" {
		host, err := s.HostIP(os.Args[2])
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(host)
	}
}
