package engine

import (
	"errors"
	"fmt"
	"rainparser/entity"
	"rainparser/io"
	"regexp"
)

var Setup io.Config
var indexes map[string]int

func StartEngine(configFile string) error {
	s, err := io.LoadConfig(configFile)

	if err != nil {
		return err
	}

	if len(s.Config.Headers) == 0 {
		return errors.New("Configuration file is empty.")
	}

	Setup = s
	io.PrintConfig(Setup)

	return nil
}

func GetColumnID(pattern string, headers []string) (int, error) {
	idx := -1

	_, err := regexp.Compile(pattern)
	if err == nil {
		for i, h := range headers {
			matched, er := regexp.MatchString(pattern, h)
			if er != nil {
				fmt.Println(er)
			} else {
				if matched {
					idx = i
				}
			}
		}
	}

	return idx, err
}

func CheckInput(headers []string) (bool, error) {
	if len(Setup.Config.Headers) == 0 {
		return false, errors.New("Invalid configuration file or there are no header info.")
	}

	if len(Setup.Config.OutFormat) == 0 {
		return false, errors.New("Invalid configuration file or there are no out order info.")
	}

	var idx int
	var err error = nil
	aux := make(map[string]int)

	for key, pattern := range Setup.Config.Headers {
		if pattern == "" {
			return false, errors.New("Invalid property:" + key)
		}

		idx, err = GetColumnID(pattern, headers)
		if err != nil || idx < 0 {
			return false, err
		} else {
			aux[key] = idx
		}
	}

	indexes = aux

	io.PrintStatus("CheckInput succeeded")
	return true, nil
}

func BuildStandard(input io.CsvFile) []entity.Employee {
	io.PrintStatus("build start")
	var employees []entity.Employee

	// skip the header
	for i := 1; i < len(input.Records); i++ {

		e := entity.Employee{
			Id:     input.Records[i][indexes[io.Properties.Id]],
			Name:   input.Records[i][indexes[io.Properties.Name]],
			Email:  input.Records[i][indexes[io.Properties.Email]],
			Salary: input.Records[i][indexes[io.Properties.Salary]],
		}

		employees = append(employees, e)
	}

	io.PrintStatus("Build succeeded")
	return employees
}
