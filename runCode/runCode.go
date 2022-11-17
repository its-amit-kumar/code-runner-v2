package runCode

// a map from string to function
import(
	"github.com/its-amit-kumar/code-runner-v2.git/language/cpp"
	"github.com/its-amit-kumar/code-runner-v2.git/language/python"
	"github.com/its-amit-kumar/code-runner-v2.git/language/java"
	"github.com/its-amit-kumar/code-runner-v2.git/language/javascript"
	"github.com/its-amit-kumar/code-runner-v2.git/language/c"
)

// filePath with name does not contain extension
func Run(fileNameWithPath string, codeLanguage string, timelimit int, memorylimit int, input string)(string, string, error, float64, int64){
	var mapOfLanguageToFunction = map[string]func(string, string, int, int)(string, string, error, float64, int64){
		"cpp" : cpp.Run,
		"python" : python.Run,
		"java" : java.Run,
		"javascript" : javascript.Run,
		"c":c.Run,
	}
	return mapOfLanguageToFunction[codeLanguage](fileNameWithPath, input, timelimit, memorylimit)
}

