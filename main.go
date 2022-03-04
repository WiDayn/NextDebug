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
	if cfg.Section("").Key("app_mode").String() == "development" {
		err = r.Run(":1016")
	} else {
		err = r.RunTLS(":8080", "./ssl.pem", "./ssl.key")
	}
	if err != nil {
		fmt.Println(err)
	}
}
