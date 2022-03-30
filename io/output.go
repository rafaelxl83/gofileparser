package io

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

func PrintStatus(statement string) {
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}

func SaveOutput(csvFile CsvFile) error {
	if csvFile.FileName == "" {
		return errors.New("Invalid file name")
	}

	if len(csvFile.Records) <= 0 {
		return errors.New("No such information")
	}

	f, err := os.Create(csvFile.FileName)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(csvFile.Records) // calls Flush internally

	if err != nil {
		return err
	}

	return nil
}
