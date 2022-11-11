package cpp

import(
	"os"
	"os/exec"
	"fmt"
	"encoding/json"
	"github.com/its-amit-kumar/code-runner-v2.git/RunExecutable"
)

func deleteFile(fileNameWithExtension string){
	e := os.Remove(fileNameWithExtension)
	if(e!=nil){

	}

}

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, error){
	app := "clang++";
	cmd := exec.Command(app, "-fsanitize=address" ,fileName+".cpp", "-o", fileName);
	_, err := cmd.Output();
	if err != nil{
		s, _ := json.MarshalIndent(err, "", "\t")
		fmt.Println(err)
		fmt.Print(string(s))
		return "", "", err
	}
	appAndArguments := []string{"./"+fileName}
	stdout, stderr, errorType := RunExecutable.Run(appAndArguments, 1, timelimit, memorylimit, input)
	deleteFile(fileName+".cpp")
	deleteFile(fileName)
	return stdout, stderr, errorType
}