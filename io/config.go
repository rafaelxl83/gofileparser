package io

import (
	"encoding/json"
	"io/ioutil"
	"rainparser/entity"
)

var Properties entity.Employee = entity.Employee{
	Id:     "id",
	Name:   "name",
	Email:  "email",
	Salary: "salary",
}

type Headers struct {
	Headers   map[string]string `json:"header"`
	OutFormat map[string]int    `json:"outorder"`
}

type Config struct {
	Config Headers `json:"config"`
}

func LoadConfig(configFile string) (Config, error) {
	f, err := ioutil.ReadFile(configFile)

	if err != nil {
		return Config{
			Config: Headers{nil, nil},
		}, err
	}

	data := Config{
		Config: Headers{},
	}

	err = json.Unmarshal([]byte(f), &data)

	if err != nil {
		return Config{
			Config: Headers{nil, nil},
		}, err
	}

	PrintStatus("config file load succeeded.")
	return data, nil
}

func PrintConfig(data Config) {
	PrintStatus("id: " + data.Config.Headers[Properties.Id])
	PrintStatus("name: " + data.Config.Headers[Properties.Name])
	PrintStatus("email: " + data.Config.Headers[Properties.Email])
	PrintStatus("salary: " + data.Config.Headers[Properties.Salary])
}
