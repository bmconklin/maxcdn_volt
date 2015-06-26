package main

import (
	"fmt"
	"log"
	"github.com/bmconklin/maxcdn_volt"
)

func main() {
	volt := maxcdn_volt.Connect("REPLACE_WITH_'ADDRESS:PORT'")
	defer volt.Conn.Close()

	companies := volt.QueryCompanies("SELECT * FROM companies limit 3")
	fmt.Println(companies)
	
	zones := volt.QueryZones("SELECT * FROM zones limit 3")
	fmt.Println(zones)
	
	urls := volt.QueryUrls("SELECT * FROM urls limit 3")
	fmt.Println(urls)
	
	rawlogs := volt.QueryRawLogs("SELECT * FROM rawlogs limit 3")
	fmt.Println(rawlogs)

	// unconventional queries

	type Count struct {
		Hits 	int
	}
	var c Count
	if err := volt.QueryOne("SELECT SUM(hits) as hits FROM companies WHERE company_id = 1738", &c); err != nil {
		log.Println(err)
	}
	fmt.Println(c)

	type ZoneCount struct{
		Hits 	int
		Zone_id int
	}
	var zc ZoneCount
	c2 := make([]ZoneCount, 0)
	resp, err := volt.QueryAll("SELECT zone_id, SUM(hits) as hits FROM zones GROUP BY zone_id")
	if err != nil {
		log.Println(err)
	}
	for resp.Table(0).HasNext() {
		resp.Table(0).Next(&zc)
		c2 = append(c2, zc)
	}

	fmt.Println(c2)
}
