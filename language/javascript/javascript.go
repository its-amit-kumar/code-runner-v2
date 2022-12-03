package javascript

import(
	"fmt"
	"os"
	"github.com/its-amit-kumar/code-runner-v2.git/RunExecutable"
)

func deleteFile(fileNameWithExtension string){
	e := os.Remove(fileNameWithExtension)
	if(e!=nil){

	}

}

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, float64, int64){
	appAndArguments := []string{"/bin/bash", "-c", "ulimit -d "+fmt.Sprint(memorylimit)+" -f  65 -u 16 -n 200 -l 64 -t "+fmt.Sprint(timelimit)+" && node "+fileName+".js"}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 2, timelimit, memorylimit, input)
	deleteFile(fileName+".js")
	return stdout, stderr, errorType, timeTaken, memoryTaken
}