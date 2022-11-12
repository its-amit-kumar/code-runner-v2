package runCode

// a map from string to function
import(
	"github.com/its-amit-kumar/code-runner-v2.git/language/cpp"
	"github.com/its-amit-kumar/code-runner-v2.git/language/python"
)

// filePath with name does not contain extension
func Run(fileNameWithPath string, codeLanguage string, timelimit int, memorylimit int, input string)(string, string, error, string, string){
	var mapOfLanguageToFunction = map[string]func(string, string, int, int)(string, string, error, string, string){
		"cpp" : cpp.Run,
		"python" : python.Run,
	}
	return mapOfLanguageToFunction[codeLanguage](fileNameWithPath, input, timelimit, memorylimit)
}

