package maxcdn_volt

import (
	"log"
	"time"
	"errors"
	"github.com/rbetts/voltdbgo/voltdb"
)

type Volt struct {
	Conn 	*voltdb.Conn
}

type Companies struct {
	Time_window 	time.Time
	Company_id 		int
	Bytes 			int64
	Hits 			int
	Cache_misses 	int
	Browser_chrome 	int
	Browser_firefox int
	Browser_mozilla int
	Browser_opera 	int
	Browser_safari 	int
	Europe 			int
	North_america 	int
	Asia 			int
	South_america 	int
	Oceania 		int
	Windows 		int
	Mac 			int
	Linux 			int
	Mobile 			int
	Android 		int
	Ios 			int
	Hits200 		int
	Hits204 		int
	Hits206 		int
	Hits301 		int
	Hits302 		int
	Hits303 		int
	Hits304 		int
	Hits307 		int
	Hits400 		int
	Hits401 		int
	Hits403 		int
	Hits404 		int
	Hits408 		int
	Hits409 		int
	Hits410 		int
	Hits412 		int
	Hits414 		int
	Hits416 		int
	Hits499 		int
	Hits500 		int
	Hits501 		int
	Hits502 		int
	Hits503 		int
	Hits504 		int
	Hits505 		int
	Hits2xx 		int
	Hits3xx 		int
	Hits4xx 		int
	Hits5xx 		int
	S0 				int
	S1 				int
	S2 				int
	S3 				int
	S4 				int
	S5 				int
	S6 				int
	S7 				int
	S8 				int
	S9 				int
}

type Zones struct {
	Time_window 	time.Time
	Zone_id 	 	int
	Company_id 		int
	Bytes 			int64
	Hits 			int
	Cache_misses 	int
	Browser_chrome 	int
	Browser_firefox int
	Browser_mozilla int
	Browser_opera 	int
	Browser_safari 	int
	Europe 			int
	North_america 	int
	Asia 			int
	South_america 	int
	Oceania 		int
	Windows 		int
	Mac 			int
	Linux 			int
	Mobile 			int
	Android 		int
	Ios 			int
	Hits200 		int
	Hits204 		int
	Hits206 		int
	Hits301 		int
	Hits302 		int
	Hits303 		int
	Hits304 		int
	Hits307 		int
	Hits400 		int
	Hits401 		int
	Hits403 		int
	Hits404 		int
	Hits408 		int
	Hits409 		int
	Hits410 		int
	Hits412 		int
	Hits414 		int
	Hits416 		int
	Hits499 		int
	Hits500 		int
	Hits501 		int
	Hits502 		int
	Hits503 		int
	Hits504 		int
	Hits505 		int
	Hits2xx 		int
	Hits3xx 		int
	Hits4xx 		int
	Hits5xx 		int
	S0 				int
	S1 				int
	S2 				int
	S3 				int
	S4 				int
	S5 				int
	S6 				int
	S7 				int
	S8 				int
	S9 				int
}

type Urls struct {
	Time_window 	time.Time
	Zone_id 		int
	Url 			string
	Company_id 		int
	Scheme 			string
	Hostname 		string
	Uri				string
	Bytes 			int64
	Hits 			int
	Cache_misses 	int
	Hits200 		int
	Hits204 		int
	Hits206 		int
	Hits301 		int
	Hits302 		int
	Hits303 		int
	Hits304 		int
	Hits307 		int
	Hits400 		int
	Hits401 		int
	Hits403 		int
	Hits404 		int
	Hits408 		int
	Hits409 		int
	Hits410 		int
	Hits412 		int
	Hits414 		int
	Hits416 		int
	Hits499 		int
	Hits500 		int
	Hits501 		int
	Hits502 		int
	Hits503 		int
	Hits504 		int
	Hits505 		int
	Hits2xx 		int
	Hits3xx 		int
	Hits4xx 		int
	Hits5xx 		int
}

type RawLogs struct {
	Ti			time.Time
	Batch_id	string
	Rawlog_id	int64
	By_tr		int64
	Br			string
	Sn			string
	Cy			string
	Co			string
	Ct			string
	Ip			string
	La			float64
	Lo			float64
	St			string
	Ci			int
	Es			string
	Hn			string
	Me			string
	Ot			float64
	Pi			int
	Pr			string
	Qs			string
	Rf			string
	Rt			float64
	Sc			string
	Si			int
	Ss			int
	Ui			string
	Ua			string
	Zi			int
}

// Connect to the Volt Database
// Address should be sent in the form "address:port"
func Connect(addr string) *Volt {
	volt, err := voltdb.NewConnection("", "", addr)
	if err != nil {
		log.Fatalf("Connection error %v\n", err)
	}
	if !volt.TestConnection() {
		log.Fatalf("Connection error: failed to ping VoltDB database.")
	}
	return &Volt{
		volt,
	}
}

// Query any table for one item, must pass in a struct to get populated. This 
// should be a struct defined on your end based on the expected query
// response.
func (volt *Volt) QueryOne(query string, resp interface{}) error {
	response, err := volt.Conn.Call("@AdHoc", query)
	if err != nil {
		log.Fatal(err)
	}
	if response.Table(0).HasNext() {
		response.Table(0).Next(resp)
	} else {
		return errors.New("No items found matching that query.")
	}
	return nil
}

// Query any table for all items matching the query, must pass in a 
// struct to get populated. This should be a struct defined on your 
// end based on the expected query response.
func (volt *Volt) QueryAll(query string) (*voltdb.Response, error) {
	return  volt.Conn.Call("@AdHoc", query)
}

// Query the Companies table. Use a complete query string.
// Returns a slice of records regardless of response set size.
func (volt *Volt) QueryCompanies(query string) []Companies {
	response, err := volt.Conn.Call("@AdHoc", query)
	if err != nil {
		log.Fatal(err)
	}

	var c Companies
	var resp []Companies
	for response.Table(0).HasNext() {
		response.Table(0).Next(&c)
		resp = append(resp, c)
	}
	return resp
}

// Query the Zones table. Use a complete query string.
// Returns a slice of records regardless of response set size.
func (volt *Volt) QueryZones(query string) []Zones {
	response, err := volt.Conn.Call("@AdHoc", query)
	if err != nil {
		log.Fatal(err)
	}

	var z Zones
	var resp []Zones
	for response.Table(0).HasNext() {
		response.Table(0).Next(&z)
		resp = append(resp, z)
	}
	return resp
}

// Query the URLS table. Use a complete query string.
// Returns a slice of records regardless of response set size.
func (volt *Volt) QueryUrls(query string) []Urls {
	response, err := volt.Conn.Call("@AdHoc", query)
	if err != nil {
		log.Fatal(err)
	}

	var u Urls
	var resp []Urls
	for response.Table(0).HasNext() {
		response.Table(0).Next(&u)
		resp = append(resp, u)
	}
	return resp
}

// Query the RAWLOGS table. Use a complete Query string.
// Returns a slice of records regardless of response set size.
func (volt *Volt) QueryRawLogs(query string) []RawLogs {
	response, err := volt.Conn.Call("@AdHoc", query)
	if err != nil {
		log.Fatal(err)
	}

	var l RawLogs
	var resp []RawLogs
	for response.Table(0).HasNext() {
		response.Table(0).Next(&l)
		resp = append(resp, l)
	}
	return resp
}