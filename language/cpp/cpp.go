package cpp

import(
	"os/exec"
	"fmt"
	"encoding/json"
)

func Run(fileName string, input string, timelimit int, memorylimit int)(string, string, string){
	app := "g++";
	cmd := exec.Command(app, fileName+".cpp", "-o", fileName);
	_, err := cmd.Output();
	if err != nil{
		s, _ := json.MarshalIndent(err, "", "\t")
		fmt.Println(err)
		fmt.Print(string(s))
		return "", "", ""
	}
	return "", "", ""
}