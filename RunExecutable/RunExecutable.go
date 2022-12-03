package RunExecutable

import (
	"os"
	"os/exec"
	"bytes"
	"strings"
	"context"
	"time"
	"fmt"
	"syscall"
	"strconv"
	"math/rand"
	"errors"
	"github.com/pbar1/pkill-go"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

func randSeq(n int) string{
	b := make([]rune, n)
	for i:=range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b);
}

func FileExists(filePath string) (bool, error) {
    info, err := os.Stat(filePath)
    if err == nil {
        return !info.IsDir(), nil
    }
    if errors.Is(err, os.ErrNotExist) {
        return false, nil
    }
    return false, err
}


func createAndReturnUser(userName string)(uint32, error){
	cmd := exec.Command("useradd", userName)
	_, err := cmd.Output()
	if err!=nil {
		fmt.Println("User Not created")
		fmt.Println(err)
		return 0, err
	}
	cmd1 := exec.Command("id", "-u", userName)
	output1,err1 := cmd1.Output()
	if(err1!=nil){
		fmt.Println("Not able to get UID")
		return 0, err1
	}
	// cmd2 := exec.Command("setquota", "-u", userName, "10MB", "10MB", "0", "0", "/")
	// _,err2 := cmd2.Output()
	// if(err2!=nil){
	// 	fmt.Println("Not able to set quota")
	// 	return 0, err2
	// }
	userIdInt := strings.TrimSuffix(string(output1),"\n")
	userId, err1 := strconv.ParseUint(string(userIdInt), 10, 32)
	return uint32(userId), err1
}

func Run(appAndArgument []string, length int, timelimit int, memorylimit int, input string)(string, string, error, float64, int64){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	timelimitConstrain, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timelimit*1000))
	defer cancel()
	userName := randSeq(10)
	userId, errUserId := createAndReturnUser(userName)
	cmd := exec.CommandContext(timelimitConstrain, appAndArgument[0], appAndArgument[1:]...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if(errUserId!=nil){
		return "", "errUserCreation", errUserId, 0, 0
	}
	
	fmt.Println(userId)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL
	cmd.SysProcAttr.GidMappingsEnableSetgroups = true
	cmd.SysProcAttr.Setpgid = true
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid:userId, Gid:userId}
	done := make(chan error, 1)
	outputSize := make(chan bool, 10)
	outputGoRoutine := make(chan bool, 1)
	startTime := time.Now()
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
					if(time.Since(startTime).Seconds() > (float64(timelimit) + 0.5)){
						_, errr := pkill.Pkill("sleep", syscall.SIGKILL)
						if(errr!=nil){
							fmt.Println("unalbe to kill sleep", errr)
						}
						for i:= 0; i<10; i++{
							killAllByUser := exec.Command("/bin/bash", "-c", "killall -u "+userName)
							_, errKill := killAllByUser.CombinedOutput()
							if(errKill!=nil){
								fmt.Println("Unable to killl user process", errKill, i)
							}
						}
						return
					}
			}
		}
	}()
	done<-cmd.Run()
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