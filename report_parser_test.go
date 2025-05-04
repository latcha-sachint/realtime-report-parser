package main

import (
	"testing"
	"time"
)

var validColumns = []string{"vin", "dealer_code", "created", "overall_severity", "delivery_status", "lead_id"}

var validRecords = [][]string{
	{"vin", "dealer_code", "created", "overall_severity", "delivery_status", "lead_id"},
	{"1V2RR2CA5MC531095", "422236", "2024-11-11 00:00:36.507", "Blue", "COMPLETE", "e323f4ff-9fe9-11ef-93f6-73eca7c6a986"},
	{"3VVMB7AX2RM103495", "402428", "0-11-11 00:04:18.697", "Red", "COMPLETE", "678956fb-9fea-11ef-b200-a11264b10928"},
	{"3VWCB7BU1LM002015", "f402168", "2024-11-11 00:05:46.303", "Blue", "COMPLETE", "9ba81ff2-9fea-11ef-b060-73eca7c6a986"},
}

var invalidHeaderRecords = [][]string{
	{"vin", "dealer_code", "created", "overall_severity", "delivery_statuses", "lead_ids"},
	{"1V2RR2CA5MC531095", "422236", "0-11-11 00:00:36.507", "Blue", "COMPLETE", "e323f4ff-9fe9-11ef-93f6-73eca7c6a986"},
}

// Tests if the file can be read from
func TestReadRecords_FromFile(t *testing.T) {
	var testReportFile string = "realtime_test_file.csv"
	t.Run("should read from a valid path and process", func(t *testing.T) {
		_, err := ReadRecords(testReportFile)

		// Test if the file can be read from a given path
		if err != nil {
			t.Errorf("failed to read from file")
		}
	})

	t.Run("should not read from an invalid path", func(t *testing.T) {
		records, err := ReadRecords("")

		if err == nil {
			t.Errorf("failed to read from file")
		}

		if records != nil {
			t.Errorf("file is not valid")
		}

	})
}

func TestHasAllRequiredColumns(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  bool
	}{
		{"should return true for all valid column headers", validColumns, true},
		{"should return false if there are invalid columns", []string{"invalid_column"}, false},
		{"should return false for empty columns", []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := hasAllRequiredColumns(tt.input)
			if isValid != tt.want {
				t.Errorf("got %t, want %t", isValid, tt.want)
			}
		})
	}
}

func TestParseReports(t *testing.T) {
	var rr = RealtimeReport{}

	t.Run("should not parse, if there aren't any records", func(t *testing.T) {
		t.Parallel()
		_, err := rr.ParseReport(&[][]string{{}})
		if err == nil {
			t.Errorf("no records to parse")
		}
	})

	t.Run("should have all the necessary column", func(t *testing.T) {
		t.Parallel()
		_, err := rr.ParseReport(&invalidHeaderRecords)
		if err == nil {
			t.Errorf("no records parsed")
		}
	})

	t.Run("should read and parse records", func(t *testing.T) {
		t.Parallel()
		parsedRecords, err := rr.ParseReport(&validRecords)

		if err != nil {
			t.Error("error while parsing")
		}

		if len(parsedRecords) == 0 {
			t.Errorf("no records parsed")
		}
	})

	t.Run("should omit records with invalid types", func(t *testing.T) {
		t.Parallel()
		parsedRecords, err := rr.ParseReport(&validRecords)

		if err != nil {
			t.Error("error while parsing")
		}

		if len(parsedRecords) != 1 {
			t.Errorf("more records than expected")
		}
	})

	t.Run("should match the type", func(t *testing.T) {
		t.Parallel()
		parsedRecords, err := rr.ParseReport(&validRecords)

		if err != nil {
			t.Error("error while parsing")
		}

		var expectedVIN string = "1V2RR2CA5MC531095"
		if parsedRecords[0].vin != expectedVIN {
			t.Errorf("type mismatch for vin, expected a string")
		}

		var expectedDealerCode uint = 422236
		if parsedRecords[0].dealerCode != expectedDealerCode {
			t.Errorf("type mismatch for dealerCode, expected an int")
		}

		expectedCreated, _ := time.Parse(time.DateTime, "2024-11-11 00:00:36.507")
		if parsedRecords[0].created != expectedCreated {
			t.Errorf("type mismatch for created, expected time.Time")
		}

		overallSeverity := "Blue"
		if parsedRecords[0].overallSeverity != overallSeverity {
			t.Errorf("type mismatch for overallSeverity, expected string")
		}

		deliveryStatus := "COMPLETE"
		if parsedRecords[0].deliveryStatus != deliveryStatus {
			t.Errorf("type mismatch for deliveryStatus, expected string")
		}

		leadId := "e323f4ff-9fe9-11ef-93f6-73eca7c6a986"
		if parsedRecords[0].leadId != leadId {
			t.Errorf("type mismatch for leadId, expected string")
		}
	})
}
