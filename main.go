package main

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
func main(){
	pathToCodeFiles,_ := os.Getwd()
	pathToCodeFiles+="/"
	mapOfExtension := map[string]string{
		"cpp" : ".cpp",
	}
	var code, codeLanguage, input string;
	var timeLimit, memoryLimit int;
	//fmt.Scanln(&code);
	code1, _ := os.ReadFile("sample-files/file.cpp")
	code = string(code1)
	fmt.Scanln(&codeLanguage);
	fmt.Scanln(&input);
	fmt.Scan(&timeLimit);
	fmt.Scan(&memoryLimit);
	//var stdout, stderr, errStatus string
	rand.Seed(time.Now().UnixNano())
	fileName := randSeq(10)
	_, err := createFile(pathToCodeFiles+fileName+mapOfExtension[codeLanguage], code)
	if err != nil{
		fmt.Println("Not of")
	}
	fmt.Println("Done")

	stdout, stderr, errStatus := runCode.Run(fileName, codeLanguage, timeLimit, memoryLimit, input)
	fmt.Println(stdout)
	fmt.Println(stderr)
	fmt.Println(errStatus)

	


}