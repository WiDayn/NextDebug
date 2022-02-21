package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"prmlk.com/nextdebug/common"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
