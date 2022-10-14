package main

import (
	"fmt"
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	//引入数据库
	model.InitDb()
	//引入路由组件
	routes.InitRouter()
	fmt.Println("test")
}
