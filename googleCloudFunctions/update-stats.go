package p

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	// "google.golang.org/genproto/googleapis/cloud/bigquery/v2"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func UpdateStats(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Uid       string `json:"uid"`
		SrvUid    string `json:"srv_uid"`
		PrimUid   string `json:"prim_uid"`
		Crn       string `json:"crn"`
		time      civil.DateTime
		timestamp civil.Time
	}

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "pc-universal-print-connector")
	if err != nil {
		fmt.Printf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	inserter := client.Dataset("connector_installation").Table("connector_monitor").Inserter()

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Hello World!")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
	now := time.Now()

	d.time = civil.DateTimeOf(now)
	d.timestamp = civil.TimeOf(now)

	if err := inserter.Put(ctx, d); err != nil {
		fmt.Printf("Failed to insert data to table: %v", err)
		return
	}
}

// if d.Message == "" {
// 	fmt.Fprint(w, "Hello World!")
// 	return
// }
// fmt.Fprint(w, html.EscapeString(d.Message))
// }
