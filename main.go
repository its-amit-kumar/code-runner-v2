package main

import (
	//"os/exec"
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "github.com/its-amit-kumar/code-runner-v2.git/createCodeSubmission"
)

type SubmitCode struct{
	Id string `json:"id"`
	Code string `json:"code"`
	TimeLimit int `json:"timeLimit"`
	MemoryLimit int `json:"memoryLimit"`
	Language string `json:"language"`
	Input string `json:"input"`
}

func SubmitCodeSubmission(c *gin.Context){
	var reqInput SubmitCode
	if err := c.BindJSON(&reqInput); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": err.Error()})
		return
	}
	stdout, stderr, err, timeTaken, memoryTaken := createCodeSubmission.CreateSubmission(reqInput.Code, reqInput.Language, reqInput.Input, reqInput.TimeLimit, reqInput.MemoryLimit)
	var errStatus string
	if(err == nil){
		errStatus = ""
	}else{
		errStatus = err.Error()
	}
	fmt.Println(errStatus)
	
	c.JSON(http.StatusOK, gin.H{
		"success":"ok",
		"stdout":stdout,
		"stderr":stderr,
		"err":errStatus,
		"timeTaken":timeTaken,
		"memoryTaken":memoryTaken,})
}

func main() {
   r := gin.Default()
   r.Use(cors.New(cors.Config{
         AllowOrigins: []string{"*"},
         AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
         AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
     }))
   r.POST("/submitCode", SubmitCodeSubmission)
   r.Run(":5300") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func main() {
// 	cmd := exec.Command("adduser", "test4")
// 	_, err := cmd.Output()
// 	if err!=nil {
// 		fmt.Println("Error occured")
// 		return
// 	}
// 	fmt.Println("User Created, getting UID and GID of the user")
// 	cmd1 := exec.Command("id", "-u", "test4")
// 	output1,err1 := cmd1.Output()
// 	if(err1!=nil){
// 		fmt.Println("Couldn't get uid")
// 		return
// 	}
// 	fmt.Printf("var1 = %T\n", output1)
// 	fmt.Println("The uid is", string(output1));
// }