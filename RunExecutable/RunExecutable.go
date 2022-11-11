package RunExecutable

import (
	"os/exec"
	"bytes"
	"strings"
	"context"
	"time"
)

func Run(appAndArgument []string, length int, timelimit int, memorylimit int, input string)(string, string, error){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	timelimitConstrain, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timelimit*1000))
	defer cancel()
	cmd := exec.CommandContext(timelimitConstrain, appAndArgument[0], appAndArgument[1:]...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	done := make(chan error, 1)
	done <- cmd.Run()
	errTLE := <-done
	if(errTLE!=nil){
		return stdout.String(), "TLE", errTLE
	}
	return stdout.String(), stderr.String(), errTLE

}