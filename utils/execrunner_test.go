package utils

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestGetCommandExits(t *testing.T) {
	command := exec.Command("uname")
	runner := CreateExecRunner(command, "/")
	output := string(runner.Run())
	fmt.Println(output)
	if len(output) < 2 {
		t.Error(output)
	}
}
