package test

import (
	_ "fmt" // no more error
	"rainparser/controller"
	"testing"
)

func TestController(t *testing.T) {
	err := controller.Run("..\\sample\\config.json", "..\\sample\\roster1.csv", "..\\sample\\roster1_parsed.csv")
	AssertEqual(t, err, nil)
}
