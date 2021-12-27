// DATABASE QUERIES/LOGICS GO HERE
package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-todo-app/config"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo) (err error) {
	if err = config.DB.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(todo *Todo) (err error) {
	if err = config.DB.Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
