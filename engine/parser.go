package engine

import (
	"errors"
	"net/mail"
	"rainparser/entity"
	"rainparser/io"
)

var emailMap map[string]bool = make(map[string]bool)
var idMap map[string]bool = make(map[string]bool)

func validMail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func CheckUniqueness(email string, id string) bool {

	if !validMail(email) {
		io.PrintStatus("Email is invalid: " + email)
		return false
	}
	if _, ok := emailMap[email]; ok {
		io.PrintStatus("Unique value violation: " + email)
		return false
	} else {
		emailMap[email] = true
	}

	if _, ok := idMap[id]; ok {
		io.PrintStatus("Unique value violation: " + id)
		return false
	} else {
		idMap[id] = true
	}

	return true
}

func buildLine(id string, name string, email string, salary string) []string {
	line := []string{"", "", "", ""}

	line[Setup.Config.OutFormat[io.Properties.Id]] = id
	line[Setup.Config.OutFormat[io.Properties.Name]] = name
	line[Setup.Config.OutFormat[io.Properties.Email]] = email
	line[Setup.Config.OutFormat[io.Properties.Salary]] = salary

	return line
}

func ParseLine(employee entity.Employee) ([]string, bool) {
	if CheckUniqueness(employee.Email, employee.Id) {
		line := buildLine(
			employee.Id,
			employee.Name,
			employee.Email,
			employee.Salary)

		io.PrintStatus("line parsed: ID=" + line[0])
		return line, true
	}

	return nil, false
}

func ParseFile(employees []entity.Employee) (io.CsvFile, error) {
	io.PrintStatus("File Parsing started")
	if len(Setup.Config.Headers) == 0 {
		return io.CsvFile{}, errors.New("Invalid configuration file or there are no data.")
	}

	csv := io.CsvFile{
		FileName: "",
		Records:  [][]string{},
	}

	header := buildLine(
		io.Properties.Id,
		io.Properties.Name,
		io.Properties.Email,
		io.Properties.Salary)
	csv.Records = append(csv.Records, header)

	for _, e := range employees {
		if s, ok := ParseLine(e); ok {
			csv.Records = append(csv.Records, s)
		} else {
			return io.CsvFile{}, errors.New("Unique value violation.")
		}
	}

	io.PrintStatus("File parsing succeeded")
	return csv, nil
}
