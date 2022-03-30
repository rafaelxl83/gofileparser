package test

import (
	_ "fmt" // no more error
	"rainparser/io"
	"testing"
)

func TestOutputEmpty(t *testing.T) {
	out := io.CsvFile{
		FileName: "",
		Records:  nil,
	}
	err := io.SaveOutput(out)
	AssertNotEqual(t, err, nil)
}

func TestOutputValidFile(t *testing.T) {
	out := io.CsvFile{
		FileName: "testout.csv",
		Records: [][]string{
			{"first_name", "last_name", "occupation"},
			{"John", "Doe", "gardener"},
			{"Lucy", "Smith", "teacher"},
			{"Brian", "Bethamy", "programmer"},
		},
	}
	err := io.SaveOutput(out)
	AssertEqual(t, err, nil)
}
