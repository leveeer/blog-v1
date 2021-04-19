package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/getArticleList", getArticleList)
	err := r.Run(":3000")
	if err != nil {
		fmt.Println("server run filed, err:", err)
	}

}

func getArticleList(c *gin.Context) {
	fmt.Println("ok")

}
