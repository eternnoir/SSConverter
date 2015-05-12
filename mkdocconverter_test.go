package ssconverter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const MKSITENAME = "utuse"

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

func TestBuild(t *testing.T) {
	createTestSite("./")
	converter, err := CreateMkdocsConverter(filepath.Join("./", MKSITENAME))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	result, builderr := converter.BuildSite()
	if builderr != nil {
		t.Error(err)
		t.Fail()
	}
	if !result {
		t.Fail()
	}
}

func createTestSite(path string) {
	os.RemoveAll(filepath.Join(path, MKSITENAME))
	command := exec.Command("mkdocs", "new", MKSITENAME)
	command.Dir = path
	output, _ := command.Output()
	fmt.Println(string(output))
}
