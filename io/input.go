package io

import (
	"encoding/csv"
	"os"
)

func LoadInput(fileName string) (CsvFile, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return CsvFile{
			FileName: "",
			Records:  [][]string{},
		}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		return CsvFile{
			FileName: "",
			Records:  [][]string{},
		}, err
	}

	return CsvFile{
		FileName: fileName,
		Records:  records,
	}, nil
}
