package ssconverter

import (
	"bytes"
	"errors"
	"github.com/eternnoir/archiveutil"
	"github.com/eternnoir/ssconverter/utils"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Mkdocs's converter struct
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

// If mkdocs command can found will return true.
func (converter *MkdocsConverter) CheckCommand() bool {
	cmd := exec.Command("mkdocs")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output := string(runner.Run())
	return strings.Contains(output, "version")
}

// Buid site from source path. If build success will return true,fail return false.
func (converter *MkdocsConverter) BuildSite() (result bool, err error) {
	if !converter.CheckCommand() {
		converter.Logger.Println("Mkdocs package not found")
		return false, errors.New("Mkdocs Not Found.")
	}
	pathExist, err := utils.CheckPathExist(converter.sourcePath)
	if err != nil {
		return false, err
	}
	if !pathExist {
		converter.Logger.Println("Mkdocs site path error.")
		return false, errors.New("Mkdocs site path error.")
	}
	cmd := exec.Command("mkdocs", "build", "--clean")
	runner := utils.CreateExecRunner(cmd, converter.sourcePath)
	output := string(runner.Run())
	converter.Logger.Println(output)
	if !strings.Contains(output, "Building documentation to directory") {
		return false, errors.New(output)
	}
	converter.Logger.Println("Build site at:" + filepath.Join(converter.sourcePath, "site"))
	return true, nil
}

// Get generated site archive bytes buffer. archiveType can be only "zip" now.
func (converter *MkdocsConverter) GetSiteBytes(archiveType string) (buffer *bytes.Buffer, err error) {
	converter.BuildSite()
	buf := new(bytes.Buffer)
	z := archiveutil.CreateArchive(archiveType, buf)
	err = z.AddFolder(filepath.Join(converter.sourcePath, "site"))
	if err != nil {
		converter.Logger.Println(err)
		return nil, err
	}
	err = z.Close()
	if err != nil {
		converter.Logger.Println(err)
		return nil, err
	}
	return buf, err
}