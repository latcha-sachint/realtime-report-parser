package main

import (
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "read_csv")
	filePath := "/Users/sachint/Desktop/realtime_report_04_28.csv"
	records, _ := ReadRecords(filePath)
	var r RealtimeReport
	r.ParseReport(records)
}
