package java

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

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error, float64,int64){
	app:="javac"
	var compileStderr, compileStdout bytes.Buffer;
	cmd := exec.Command(app, "Main.java");
	cmd.Stdout = &compileStdout
	cmd.Stderr = &compileStderr
	err := cmd.Run()
	if err != nil{
		deleteFile("Main.java")
		deleteFile("Main.class")
		return compileStdout.String(), compileStderr.String(), err, float64(0), int64(0)
	}
	appAndArguments := []string{"java", "Main"}
	stdout, stderr, errorType, timeTaken, memoryTaken := RunExecutable.Run(appAndArguments, 2, timelimit, memorylimit, input)
	deleteFile("Main.java")
	deleteFile("Main.class")
	return stdout, stderr, errorType, timeTaken, memoryTaken
}
