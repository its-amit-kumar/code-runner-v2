package python

import(
	"os"
	"github.com/its-amit-kumar/code-runner-v2.git/RunExecutable"
)

func deleteFile(fileNameWithExtension string){
	e := os.Remove(fileNameWithExtension)
	if(e!=nil){

	}

}

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, string, string){
	appAndArguments := []string{"python3", fileName+".py"}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 2, timelimit, memorylimit, input)
	deleteFile(fileName+".py")
	return stdout, stderr, errorType, timeTaken, memoryTaken
}