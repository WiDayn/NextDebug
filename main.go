package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"prmlk.com/nextdebug/common"
)

func main() {
	cfg, err := ini.Load("config/app.ini")
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	err = r.Run(":" + cfg.Section("server").Key("port").String())
	if err != nil {
		fmt.Println(err)
	}
}
