package main

import (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//container.in
	r.GET("/ping", func(c *gin.Context) {
		userBiz := InitUserBiz("dbstr")
		fmt.Println("userbizBiz:%+v", *userBiz)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and s  erve on 0.0.0.0:8080 (for windows "localhost:8080")

}
