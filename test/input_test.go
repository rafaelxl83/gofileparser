package test

import (
	"fmt"
	"rainparser/io"
	"testing"
)

func TestInputEmpty(t *testing.T) {
	in, err := io.LoadInput("test.csv")
	_ = err
	AssertEqual(t, in.FileName, "")
}

func TestInputValidFile(t *testing.T) {
	in, err := io.LoadInput("../sample/roster1.csv")
	_ = err
	AssertNotEqual(t, in.FileName, "")
}

func TestInputRecords(t *testing.T) {
	in, err := io.LoadInput("../sample/roster1.csv")
	fmt.Println(err)
	AssertNotEqual(t, len(in.Records), 0)
}
