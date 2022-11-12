package RunExecutable

import (
	"os/exec"
	"bytes"
	"strings"
	"context"
	"time"
	//"fmt"
	"syscall"
	"strconv"
)



func Run(appAndArgument []string, length int, timelimit int, memorylimit int, input string)(string, string, error, string, string){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	timelimitConstrain, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timelimit*1000))
	defer cancel()
	cmd := exec.CommandContext(timelimitConstrain, appAndArgument[0], appAndArgument[1:]...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	done := make(chan error, 1)
	outputSize := make(chan bool, 10)
	startTime := time.Now()
	done <- cmd.Run()
	go func(){
		for{
			if stdout.Len() >= 65536{
				outputSize <- true
			}
		}
	}()
	select{
	case errTLE := <-done:
		timeElapsed := time.Since(startTime)
		memoryConsumed := cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
		memoryConsumedString := strconv.FormatInt(memoryConsumed, 10)
		if(errTLE!=nil){
			//fmt.Println("Killing Code", errTLE)
			//fmt.Println("killing code", stderr.String())
			if(errTLE.Error() == "signal: killed"){
				return stdout.String(), "TLE", errTLE, timeElapsed.String(), memoryConsumedString
			}
			//fmt.Println(errTLE.Error())
			return stdout.String(), stderr.String(), errTLE, timeElapsed.String(), memoryConsumedString
		}
		if(int(memoryConsumed)>memorylimit){
			//fmt.Println(memoryConsumed)
			return "", "kiledMem", errTLE, timeElapsed.String(), memoryConsumedString
		}
		return stdout.String(), stderr.String(), errTLE, timeElapsed.String(), memoryConsumedString
	case <- outputSize:
		timeElapsed := time.Since(startTime)
		memoryConsumed := cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
		memoryConsumedString := strconv.FormatInt(memoryConsumed, 10)
		return "", "KilledOutput", nil , timeElapsed.String(), memoryConsumedString
	}

}