package cpp

import(
	"os/exec"
	"fmt"
)

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, string){
	app := "g++";
	cmd := exec.Command(app, fileName+".cpp", "-o", fileName);
	_, err := cmd.Output();
	if err != nil{
		fmt.Printf("%+v\n", err);
		return "", "", ""
	}
	return "", "", ""
}