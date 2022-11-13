package createCodeSubmission

import(
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/its-amit-kumar/code-runner-v2.git/runCode"
)
/*
4 inputs
input file name
language
time limit
memory limit
*/

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

func randSeq(n int) string{
	b := make([]rune, n)
	for i:=range b{
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b);
}

func createFile(pathToCodeFileWithName string, code string) (string, error){
	f, err := os.Create(pathToCodeFileWithName)
	if err != nil{
		return code, err;
	}
	defer f.Close()

	_, err2 := f.WriteString(code)
	
	if err2!=nil{
		return code, err
	}
	return code, nil;

}


func CreateSubmission(code string, codeLanguage string, input string, timeLimit int, memoryLimit int)(string, string, error, float64, int64) {
	pathToCodeFiles,_ := os.Getwd()
	pathToCodeFiles+="/"
	mapOfExtension := map[string]string{
		"cpp" : ".cpp",
		"python" : ".py",
		"java" : ".java",
		"javascript" : ".js",
	}
	rand.Seed(time.Now().UnixNano())
	fileName := randSeq(10)
	if(codeLanguage == "java"){
		fileName = "Main"
	}
	_, err := createFile(pathToCodeFiles+fileName+mapOfExtension[codeLanguage], code)
	if err != nil{
		fmt.Println(err);
	}
	stdout, stderr, errStatus, timeTaken, memoryTaken := runCode.Run(fileName, codeLanguage, timeLimit, memoryLimit, input)
	return stdout, stderr, errStatus, timeTaken, memoryTaken

	


}