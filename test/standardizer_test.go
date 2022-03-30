package test

import (
	_ "fmt" // no more error
	"rainparser/engine"
	"rainparser/io"
	"testing"
)

func TestStandardizerStartFail(t *testing.T) {
	err := engine.StartEngine("config.json")
	AssertNotEqual(t, err, nil)
}

func TestStandardizerStartOk(t *testing.T) {
	err := engine.StartEngine("../sample/config.json")
	AssertEqual(t, err, nil)
}

func TestStandardizerGetColumnID(t *testing.T) {
	idx, err := engine.GetColumnID("(?i)name|first", []string{"First Name", "Email", "Wage", "Number"})
	_ = err
	AssertEqual(t, idx, 0)
}

func TestStandardizerGetColumnIDFail(t *testing.T) {
	idx, err := engine.GetColumnID("(?i)name|first", []string{"Flat", "Email", "Wage", "Number"})
	_ = err
	AssertEqual(t, idx, -1)
}

func TestStandardizerCheckInputOK(t *testing.T) {
	err := engine.StartEngine("../sample/config.json")
	isOk, err := engine.CheckInput([]string{"First Name", "Email", "Wage", "Number"})
	_ = err
	AssertEqual(t, isOk, true)
}

func TestStandardizerCheckInputFail(t *testing.T) {
	err := engine.StartEngine("../sample/config.json")
	isOk, err := engine.CheckInput([]string{"Flag", "E-mail", "Salary", "ID"})
	_ = err
	AssertEqual(t, isOk, false)
}

func TestStandardizerBuild(t *testing.T) {
	err := engine.StartEngine("../sample/config.json")
	in, err := io.LoadInput("../sample/roster1.csv")
	isOk, err := engine.CheckInput(in.Records[0])

	if isOk {
		employees := engine.BuildStandard(in)
		_ = err
		AssertNotEqual(t, employees, nil)
	} else {
		AssertFail(t, in, err)
	}
}
