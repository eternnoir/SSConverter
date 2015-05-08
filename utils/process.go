package utils

import (
	"os/exec"
)

type ExecRunner struct {
	command *exec.Cmd
	dir     string
}

func CreateExecRunner(command *exec.Cmd, dir string) *ExecRunner {
	pro := new(ExecRunner)
	pro.command = command
	pro.dir = dir
	return pro
}

func (runner *ExecRunner) Run() {

}
