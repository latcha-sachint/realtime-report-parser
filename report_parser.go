package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type RealtimeReport struct {
	vin             string
	dealerCode      uint
	created         time.Time
	overallSeverity string
	deliveryStatus  string
	leadId          string
	isValid         bool
}

type ReportColumnHead struct {
	index uint8
	label string
}

// Mapping of CSV column header positions to their respective labels.
// Not all column fields/labels are currently used in the program, but they might be useful in the future.
// If a column is shifted or removed, updating this mapping is simpler than modifying the logic elsewhere.
var rc = map[string]ReportColumnHead{
	"vin":              {index: 0, label: "vin"},
	"dealer_code":      {index: 1, label: "dealer_code"},
	"created":          {index: 2, label: "created"},
	"overall_severity": {index: 3, label: "overall_severity"},
	"delivery_status":  {index: 4, label: "delivery_status"},
	"lead_id":          {index: 5, label: "lead_id"},
}

// Validates that all required headers exist in the provided data.
func hasAllRequiredColumns(headers []string) bool {
	for h := range rc {
		if !slices.Contains(headers, h) {
			return false
		}
	}
	return true
}

// ParseReport parses records which is a nested array and returns the structed RealtimeReport.
// First line of the Record is the
func (r *RealtimeReport) ParseReport(records *[][]string) ([]RealtimeReport, error) {
	headers := (*records)[0]

	// Check if there are enough Records to work with
	if len(*records) < 1 {
		return nil, errors.New("no records")
	}
	// Check if all the required column headers exist
	if !hasAllRequiredColumns(headers) {
		return nil, errors.New("doesn't have all valid column headers")
	}

	realtimeReports := make([]RealtimeReport, 0, len(*records))

	for _, row := range (*records)[1:] {
		var reportRecord RealtimeReport = RealtimeReport{isValid: true}
		for i, value := range row {
			switch headers[i] {
			case "vin":
				reportRecord.vin = value
			case "created":
				parsedDateTime, err := time.Parse(time.DateTime, value)
				if err != nil {
					fmt.Printf("skipping row with VIN %s - error while parsing created date\n", reportRecord.vin)
					reportRecord.isValid = false
					continue
				}
				// Set parsed `time.Time` Created Datetime to Report Record
				reportRecord.created = parsedDateTime
			case "dealer_code":
				dealerCode, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("skipping row with VIN %s - error while parsing dealer_code\n", reportRecord.vin)
					reportRecord.isValid = false
					continue
				}
				// Set parsed `int` DealerCode to Report Record
				reportRecord.dealerCode = uint(dealerCode)
			case "delivery_status":
				reportRecord.deliveryStatus = value
			case "overall_severity":
				reportRecord.overallSeverity = value
			case "lead_id":
				reportRecord.leadId = value
			default:
			}
		}

		// append record only if it is valid
		if reportRecord.isValid {
			realtimeReports = append(realtimeReports, reportRecord)
		}
	}

	return realtimeReports, nil
}

// Opens the specified file, reads its contents as CSV, and returns all records.
// Logs a fatal error if the file cannot be opened or the contents cannot be read.
func ReadRecords(filePath string) (*[][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer f.Close()

	// Create a new CSV reader
	csvReader := csv.NewReader(f)

	// Read all records
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &records, nil
}
