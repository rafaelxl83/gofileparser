package main

import (
	"flag"
	"io/ioutil"
	"rainparser/controller"
	"rainparser/io"
	"strings"
)

func createBadFile(fileName string) error {
	bytesRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	badfile := strings.Replace(fileName, ".csv", "_bad.csv", 1)
	err = ioutil.WriteFile(badfile, bytesRead, 0755)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	inputFile := flag.String("input", "sample\\roster1.csv", "an input file")
	outputFile := flag.String("output", "sample\\roster1_parsed.csv", "an output file")
	configFile := flag.String("config", "sample\\config.json", "a config file")
	flag.Parse()

	io.PrintStatus("input: " + *inputFile)
	io.PrintStatus("output: " + *outputFile)
	io.PrintStatus("config:" + *configFile)

	err := controller.Run(*configFile, *inputFile, *outputFile)
	if err != nil {
		io.PrintStatus(err.Error())
		createBadFile(*inputFile)
	} else {
		io.PrintStatus("Parse action completed")
	}
}
