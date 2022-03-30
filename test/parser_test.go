package test

import (
	"fmt"
	"rainparser/engine"
	"rainparser/entity"
	"testing"
)

func TestParserCheck(t *testing.T) {
	ok := engine.CheckUniqueness("email@test.com", "RT21")
	AssertEqual(t, ok, true)
}

func TestParserCheckFail(t *testing.T) {
	ok := engine.CheckUniqueness("email@test.com", "RT21")
	ok = engine.CheckUniqueness("emailtest.com", "RT21")
	ok = engine.CheckUniqueness("email@test.com", "RT21")
	ok = engine.CheckUniqueness("email2@test.com", "RT21")
	AssertEqual(t, ok, false)
}

func TestParserLine(t *testing.T) {
	employee := entity.Employee{
		Id:     "RT21",
		Name:   "Jhon Doe",
		Email:  "email@test.com",
		Salary: "$10",
	}

	err := engine.StartEngine("../sample/config.json")
	_ = err
	s, ok := engine.ParseLine(employee)
	fmt.Println(ok)
	AssertNotEqual(t, s, "")
}

func TestParserAll(t *testing.T) {
	employees := []entity.Employee{
		{
			Name:   "Jhon Doe",
			Email:  "email@test.com",
			Salary: "$10",
			Id:     "RT21",
		},
		{
			Name:   "Aloe Vera",
			Email:  "email2@test.com",
			Salary: "$4000000",
			Id:     "RT11",
		},
		{
			Name:   "Jhon Carter",
			Email:  "email3@test.com",
			Salary: "$666",
			Id:     "RT25",
		},
		{
			Name:   "Jhoana Dark",
			Email:  "email4@test.com",
			Salary: "$123124",
			Id:     "RT01",
		},
	}

	err := engine.StartEngine("../sample/config.json")
	csv, err := engine.ParseFile(employees)
	fmt.Println(csv.FileName)
	AssertEqual(t, err, nil)
}
