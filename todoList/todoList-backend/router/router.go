package router

import (
	"github.com/gin-gonic/gin"
	"todoList/controller"
)

func SetupRouter() *gin.Engine {
	r :=gin.Default()
	// find template
	r.LoadHTMLGlob("templates/*")

	// find static
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// C
		v1Group.POST("/todo", controller.CreateTodo)
		// R
		// 查看所有
		v1Group.GET("/todo", controller.GetTodoList)

		// U
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// D
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}