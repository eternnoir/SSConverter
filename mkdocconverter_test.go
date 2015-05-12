package ssconverter

import (
	"testing"
)

func TestCreateMkConverter(t *testing.T) {
	_, err := CreateMkdocsConverter("./")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateMkConverterPathNotExits(t *testing.T) {
	_, err := CreateMkdocsConverter("./fasdfasdfas")
	if err != nil {
		t.Log(err)
	} else {
		t.Fail()
	}
}
