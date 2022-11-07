package main

import(
	"fmt"
	"math/rand"
	"os"
)
/*
4 inputs
input file name
language
time limit
memory limit
*/

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");
var pathToCodeFiles = "E:/code-runner-backend/"
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
	var code, codeLanguage, input string;
	var timeLimit, memoryLimit int;
	fmt.Scanln(&code);
	fmt.Scanln(&codeLanguage);
	fmt.Scanln(&input);
	fmt.Scan(&timeLimit);
	fmt.Scan(&memoryLimit);
	//var stdout, stderr, errStatus string
	fileName := randSeq(10)
	_, err := createFile(pathToCodeFiles+fileName, code)
	if err != nil{
		fmt.Println("Not of")
	}
	fmt.Println("Done")

	//stdout, stderr, errStatus = runCode.run(fileName, codeLanguage, timeLimit, memoryLimit);

	


}