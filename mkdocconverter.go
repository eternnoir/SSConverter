package ssconverter

import (
	"errors"
	"github.com/eternnoir/ssconverter/utils"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type MkdocsConverter struct {
	sourcePath string
	Logger     *log.Logger
}

// Create a mkdoc converter.Source path is your mkdocs site doc path. If source path not
// existed or mkdoc not installed will return error.
func CreateMkdocsConverter(sourcePath string) (converter *MkdocsConverter, err error) {
	ret := new(MkdocsConverter)
	absPath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}
	ret.sourcePath = absPath
	if !ret.CheckCommand() {
		return nil, errors.New("Mkdoc command not found.")
	}
	ret.Logger = log.New(os.Stdout,
		"SSCONVERTER: ",
		log.Ldate|log.Ltime)
	return ret, nil
}

func (converter *MkdocsConverter) CheckCommand() bool {
	cmd := exec.Command("mkdocs")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output := string(runner.Run())
	return strings.Contains(output, "json")
}

// Buid site from source path. It will return command output.
func (converter *MkdocsConverter) BuildSite() (output string, err error) {
	if !converter.CheckCommand() {
		converter.Logger.Println("Mkdocs package not found")
		return "", errors.New("Mkdocs Not Found.")
	}
	pathExist, err := utils.CheckPathExist(converter.sourcePath)
	if err != nil {
		return "", err
	}
	if !pathExist {
		converter.Logger.Println("Mkdocs site path error.")
		return "", errors.New("Mkdocs site path error.")
	}
	cmd := exec.Command("mkdocs", "build", "--clean")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output = string(runner.Run())
	converter.Logger.Println(output)
	sitePath := filepath.Join(converter.sourcePath, "site")
	converter.Logger.Println("Build site at:" + sitePath)
	return output, nil
}