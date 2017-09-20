package config

import (
	"testing"
)

type DataTest struct {
	Line int
}

func TestGlobalConfig(t *testing.T) {
	Load([]byte(`{
		"test": true,
		"hello": 1.0,
		"world": "hi",
		"data": {
			"line":12
		}
	}`))

	if Has("test1") {
		t.Errorf("check not exists config item -- failed")
	}

	value, ok := Get("data.line").(float64)

	if !ok {
		t.Errorf("check data.line")
	}

	if value != 12 {
		t.Errorf("data.line != 12")
	}

	value2 := GetString("world", "test")

	if !Has("world") {
		t.Errorf("check data.line")
	}

	if value2 != "hi" {
		t.Errorf("check world -- failed")
	}

	var test DataTest

	err := GetObject("data", &test)

	if err != nil {
		t.Error(err)
	}

	println(test.Line)
}
