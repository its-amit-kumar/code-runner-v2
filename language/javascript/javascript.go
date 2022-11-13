package javascript

import(
	"os"
	"github.com/its-amit-kumar/code-runner-v2.git/RunExecutable"
)

func deleteFile(fileNameWithExtension string){
	e := os.Remove(fileNameWithExtension)
	if(e!=nil){

	}

}

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, float64, int64){
	appAndArguments := []string{"node", fileName+".js"}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 2, timelimit, memorylimit, input)
	deleteFile(fileName+".js")
	return stdout, stderr, errorType, timeTaken, memoryTaken
}