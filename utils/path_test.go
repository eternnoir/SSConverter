package utils

import (
	"testing"
)

func TestPathCheck(t *testing.T) {
	result, err := CheckPathExist("./")
	if err != nil {
		t.Error(err)
	}
	if result == false {
		t.Fail()
	}
}

func TestPathCheckNoPath(t *testing.T) {
	result, err := CheckPathExist("./ggg")
	if err != nil {
		t.Error(err)
	}
	if result == true {
		t.Fail()
	}
}
