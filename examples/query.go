package main

import (
	"fmt"
	"github.com/bmconklin/maxcdn_volt"
)

func main() {
	volt := maxcdn_volt.Connect("PLACE_ADDRESS:PORT_HERE")
	defer volt.Conn.Close()

	companies := volt.QueryCompanies("SELECT * FROM companies limit 3")
	fmt.Println(companies)
	
	zones := volt.QueryZones("SELECT * FROM zones limit 3")
	fmt.Println(zones)
	
	urls := volt.QueryUrls("SELECT * FROM urls limit 3")
	fmt.Println(urls)
	
	rawlogs := volt.QueryRawLogs("SELECT * FROM rawlogs limit 3")
	fmt.Println(rawlogs)
}
