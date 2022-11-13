package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
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
  r.POST("/submitCode", SubmitCodeSubmission)
  r.Run(":5300") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}