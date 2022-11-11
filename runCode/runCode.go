package runCode

// a map from string to function
import(
	"github.com/its-amit-kumar/code-runner-v2.git/language/cpp"
)

// filePath with name does not contain extension
func Run(fileNameWithPath string, codeLanguage string, timelimit int, memorylimit int, input string)(string, string, error){
	var mapOfLanguageToFunction = map[string]func(string, string, int, int)(string, string, error){
		"cpp" : cpp.Run,
	}
	return mapOfLanguageToFunction[codeLanguage](fileNameWithPath, input, timelimit, memorylimit)
}

