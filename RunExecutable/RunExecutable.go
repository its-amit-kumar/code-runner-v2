package RunExecutable

import (
	"os/exec"
	"bytes"
	"strings"
)

func Run(appAndArgument []string, length int, timelimit int, memorylimit int, input string)(string, string, error){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	
	cmd := exec.Command(appAndArgument[0], appAndArgument[1:]...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	done := make(chan error, 1)
	done <- cmd.Run()
	err := <-done
	return stdout.String(), stderr.String(), err

}