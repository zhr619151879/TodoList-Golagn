package model

import (
	"todoList/dao"
)

// Todo model
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// Todo CRUD
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList() (todoList []Todo, err error) {
	 //dao.DB.Find(&todoList)
	if err :=dao.DB.Find(&todoList).Error;err!= nil {
		return nil, err
	}
	return
}

func GetTodoById(id string)(todo *Todo, err error)  {
	// 查询数据库 if exist
	if err = dao.DB.Where("id=?",id).First(todo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string)(err error) {
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}