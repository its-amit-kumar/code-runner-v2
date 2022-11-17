package RunExecutable

import (
	"os/exec"
	"bytes"
	"strings"
	"context"
	"time"
	//"fmt"
	"syscall"
)



func Run(appAndArgument []string, length int, timelimit int, memorylimit int, input string)(string, string, error, float64, int64){
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
	outputGoRoutine := make(chan bool, 1)
	go func(){
		for{
			select{ 
				case <- outputGoRoutine:
					return
				default:
					if(stdout.Len() > 65536){
						errr := cmd.Process.Kill()
						if(errr!=nil){

						}
						return
					}
			}
		}
	}()
	startTime := time.Now()
	done <- cmd.Run()
	select{
	case errTLE := <-done:
		outputGoRoutine<-true
		timeElapsed := time.Since(startTime).Seconds()
		memoryConsumed := cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
		if(stdout.Len()>65536){
			return "", "KilledOutput", nil , timeElapsed, memoryConsumed
		}
		if(int(memoryConsumed)>memorylimit){
			//fmt.Println(memoryConsumed)
			return "", "kiledMem", errTLE, timeElapsed, memoryConsumed
		}
		if(errTLE!=nil){
			//fmt.Println("Killing Code", errTLE)
			//fmt.Println("killing code", stderr.String())
			if(errTLE.Error() == "signal: killed"){
				return stdout.String(), "TLE", errTLE, timeElapsed, memoryConsumed
			}
			//fmt.Println(errTLE.Error())
			return stdout.String(), stderr.String(), errTLE, timeElapsed, memoryConsumed
		}
		
		return stdout.String(), stderr.String(), errTLE, timeElapsed, memoryConsumed
	case <- outputSize:
		outputGoRoutine<-true
		timeElapsed := time.Since(startTime).Seconds()
		memoryConsumed := cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
		return "", "KilledOutput", nil , timeElapsed, memoryConsumed
	}

}