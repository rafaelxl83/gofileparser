package test

import (
	_ "fmt" // no more error
	"rainparser/io"
	"testing"
)

func TestConfigLoadNoFile(t *testing.T) {
	data, err := io.LoadConfig("config.json")
	_ = err
	AssertEqual(t, len(data.Config.Headers), 0)
}

func TestConfigLoad(t *testing.T) {
	file := "../sample/config.json"
	data, err := io.LoadConfig(file)
	_ = err
	AssertNotEqual(t, data.Config.Headers["id"], nil)
}
