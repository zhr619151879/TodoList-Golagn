package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoList/model"
)

/*
	url--> controller --> service -->model
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(context *gin.Context) {
	// 取得表单数据并绑定
	var todoTmp model.Todo
	_ = context.BindJSON(&todoTmp)
	// insert into db
	err := model.CreateATodo(&todoTmp)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		// 创建成功, 数据返回给前端
		context.JSON(http.StatusOK, todoTmp)
		//context.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todoTmp,
		//})
	}
}

func GetTodoList(context *gin.Context) {
	// 查询表中所有数据
	todoList, err := model.GetTodoList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		context.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id!"})
	}

	todo, err := model.GetTodoById(id)
	if err != nil{
		return
	}
	// get data from front-end
	_ = context.BindJSON(&todo)
	//fmt.Println(todo)
	if err = model.UpdateATodo(todo); err != nil {
		context.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id!"})
	}
	// del
	if err := model.DeleteATodo(id); err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"id:": "he"})
	}
}
