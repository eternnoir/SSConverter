package ssconverter

import (
	"errors"
	"fmt"
	"github.com/eternnoir/ssconverter/utils"
	"os/exec"
	"strings"
)

type MkdocsConverter struct {
	sourcePath string
	targetPath string
}

func CreateMkdocsConverter(sourcePath, targetPath string) *MkdocsConverter {
	ret := new(MkdocsConverter)
	ret.sourcePath = sourcePath
	ret.targetPath = targetPath
	return ret
}

func (converter *MkdocsConverter) CheckCommand() bool {
	cmd := exec.Command("mkdocs")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output := string(runner.Run())
	return strings.Contains(output, "json")
}

func (converter *MkdocsConverter) BuildSite() (sitePath string, err error) {
	if !converter.CheckCommand() {
		fmt.Println("Mkdoc not found")
		return "", errors.New("Mkdocs Not Found.")
	}
	cmd := exec.Command("mkdocs", "build")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output := string(runner.Run())
	fmt.Println(output)
	return output, nil
}