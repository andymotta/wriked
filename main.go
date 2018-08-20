package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	var timesheet = flag.String("f", "", "Excel file containing hours to upload to Wrike")
	var token = flag.String("a", "", "Get Dangerous with your Wrike API Token")
	flag.Parse()
	if *timesheet == "" || *token == "" {
		flag.Usage()
		log.Fatalf("Please enter the required arguments.")
	}

	auth := fmt.Sprintf("bearer %s", *token)

	type wriked struct {
		ID          string  `json:"id"`
		Hours       float64 `json:"hours"` // not used, match potential
		TrackedDate string  `json:"trackedDate"`
	}
	type wrikedResponse struct {
		Data []wriked `json:"data"`
	}

	var wrikeds wrikedResponse

	// get Wrike User ID via token
	req, _ := http.NewRequest("GET", "https://www.wrike.com/api/v3/contacts?&me", nil)
	req.Header.Add("authorization", auth)
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("The HTTP req failed with error %s\n", err)
	}
	json.NewDecoder(res.Body).Decode(&wrikeds)
	user := wrikeds.Data[0].ID

	xlsx, err := excelize.OpenFile(*timesheet)
	if err != nil {
		log.Fatalln(err)
	}
	for index, name := range xlsx.GetSheetMap() {
		if index != 1 {
			for c := 11; c <= 17; c++ {
				// Get value from cell by given worksheet name and axis.
				hours := xlsx.GetCellValue(name, "G"+strconv.Itoa(c))
				xdate := xlsx.GetCellValue(name, "C"+strconv.Itoa(c))
				dt, _ := time.Parse("01-02-06", xdate)
				date := dt.Format("2006-01-02")
				if hours == "0.00" || hours == "" {
					continue
				} else {
					url := fmt.Sprintf("https://www.wrike.com/api/v3/contacts/%s/timelogs?trackedDate={\"start\":\"%sT00:00:00\",\"end\":\"%sT23:59:59\"}", user, date, date)
					req, _ := http.NewRequest("GET", url, nil)
					req.Header.Add("authorization", auth)
					client := &http.Client{Timeout: 10 * time.Second}
					res, err := client.Do(req)
					if err != nil {
						log.Fatalf("The HTTP req failed with error %s\n", err)
					}
					json.NewDecoder(res.Body).Decode(&wrikeds)
					if len(wrikeds.Data) == 0 {
						fmt.Printf("Adding %s Hours for %s\n", hours, date)
						req, _ := http.NewRequest("POST", "https://www.wrike.com/api/v3/tasks/IEABW3IFKQHICKWP/timelogs", nil)
						req.Header.Set("Content-Type", "application/json")
						req.Header.Add("Authorization", auth)
						data := req.URL.Query()
						data.Add("hours", hours)
						data.Add("trackedDate", date)
						data.Add("comment", *timesheet)
						req.URL.RawQuery = data.Encode()
						// fmt.Println(req.URL.String())
						client := &http.Client{Timeout: 10 * time.Second}
						_, err := client.Do(req)
						if err != nil {
							log.Fatalf("The HTTP req failed with error %s\n", err)
						}
					} else {
						fmt.Printf("Skipping existing time entry for %s\n", date)
					}
				}
			}
		}
	}
}
