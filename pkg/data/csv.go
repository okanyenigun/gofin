package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CSVRecord struct {
	Datetime    string
	Open        float64
	High        float64
	Low         float64
	Close       float64
	Volume      int64
	Dividends   float64
	StockSplits float64
}

type CSVData struct {
	Headers []string
	Records []CSVRecord
}

func LoadCSV(filename string) (*CSVData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("CSV file is empty")
	}

	headers := rows[0]

	records := make([]CSVRecord, 0, len(rows)-1)
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		record, err := parseCSVRecord(row)
		if err != nil {
			continue
		}
		records = append(records, record)
	}

	return &CSVData{
		Headers: headers,
		Records: records,
	}, nil
}

func parseCSVRecord(row []string) (CSVRecord, error) {
	record := CSVRecord{
		Datetime: row[0],
	}

	var err error

	if record.Open, err = strconv.ParseFloat(row[1], 64); err != nil {
		return record, err
	}
	if record.High, err = strconv.ParseFloat(row[2], 64); err != nil {
		return record, err
	}
	if record.Low, err = strconv.ParseFloat(row[3], 64); err != nil {
		return record, err
	}
	if record.Close, err = strconv.ParseFloat(row[4], 64); err != nil {
		return record, err
	}
	if record.Volume, err = strconv.ParseInt(row[5], 10, 64); err != nil {
		return record, err
	}
	if record.Dividends, err = strconv.ParseFloat(row[6], 64); err != nil {
		return record, err
	}
	if record.StockSplits, err = strconv.ParseFloat(row[7], 64); err != nil {
		return record, err
	}

	return record, nil
}

func (d *CSVData) GetClosePrices() []float64 {
	closePrices := make([]float64, len(d.Records))
	for i, record := range d.Records {
		closePrices[i] = record.Close
	}
	return closePrices
}
