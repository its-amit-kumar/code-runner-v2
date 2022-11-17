package c

import(
	"os"
	"os/exec"
//"fmt"
	"github.com/its-amit-kumar/code-runner-v2.git/RunExecutable"
	"bytes"
)

func deleteFile(fileNameWithExtension string){
	e := os.Remove(fileNameWithExtension)
	if(e!=nil){

	}

}
// add support for clang
func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, float64,int64){
	var compileStdout, compileStderr bytes.Buffer
	app := "gcc";
	cmd := exec.Command(app, "-fsanitize=address", fileName+".c", "-o", fileName);
	cmd.Stdout = &compileStdout
	cmd.Stderr = &compileStderr
	err := cmd.Run()
	if err != nil{
		deleteFile(fileName+".c")
		deleteFile(fileName)
		return compileStdout.String(), compileStderr.String(), err, float64(0), int64(0)
	}
	appAndArguments := []string{"./"+fileName}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 1, timelimit, memorylimit, input)
	deleteFile(fileName+".c")
	deleteFile(fileName)
	return stdout, stderr, errorType, timeTaken, memoryTaken
}

