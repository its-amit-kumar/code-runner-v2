package cpp

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

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, string,string){
	var compileStdout, compileStderr bytes.Buffer
	app := "g++";
	cmd := exec.Command(app, "-fsanitize=address", fileName+".cpp", "-o", fileName);
	cmd.Stdout = &compileStdout
	cmd.Stderr = &compileStderr
	err := cmd.Run()
	if err != nil{
		return compileStdout.String(), compileStdout.String(), err, "", ""
	}
	appAndArguments := []string{"./"+fileName}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 1, timelimit, memorylimit, input)
	deleteFile(fileName+".cpp")
	deleteFile(fileName)
	return stdout, stderr, errorType, timeTaken, memoryTaken
}