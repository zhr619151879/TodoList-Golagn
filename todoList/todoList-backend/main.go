package main

import (
	"todoList/dao"
	"todoList/model"
	"todoList/router"
)

func main() {
	// conn mysql db
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	defer dao.DB.Close()
	// associate Todo
	dao.DB.AutoMigrate(&model.Todo{}) // table name: Todos

	// start
	r:= router.SetupRouter()
	r.Run(":8080")
}
