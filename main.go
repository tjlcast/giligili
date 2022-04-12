package main

import (
	"giligili/conf"
	"giligili/server"
	"giligili/util"
	"os"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	port := os.Getenv("PORT")
	util.Log().Info("Listen on port: %s.", port)

	r := server.NewRouter()
	r.Run(":" + port)
}
