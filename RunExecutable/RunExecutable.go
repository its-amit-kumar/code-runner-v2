package RunExecutable

import (
	"os/exec"
	"bytes"
	"strings"
	"context"
	"time"
	"fmt"
	"syscall"
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
	memoryConsumed := cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
	if(errTLE!=nil){
		return stdout.String(), "TLE", errTLE
	}
	if(int(memoryConsumed)>memorylimit){
		fmt.Println(memoryConsumed)
		return "", "kiledMem", errTLE
	}
	return stdout.String(), stderr.String(), errTLE

}