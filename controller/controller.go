package controller

import (
	"rainparser/engine"
	"rainparser/io"
)

func Run(config string, input string, output string) error {

	err := engine.StartEngine(config)
	if err != nil {
		return err
	}

	in, err := io.LoadInput(input)
	if err != nil {
		return err
	}

	if ok, err := engine.CheckInput(in.Records[0]); !ok {
		return err
	}

	employees := engine.BuildStandard(in)

	out, err := engine.ParseFile(employees)
	if err != nil {
		return err
	}

	io.PrintStatus("Saiving output file")
	out.FileName = output
	err = io.SaveOutput(out)
	if err != nil {
		return err
	}

	return nil
}
