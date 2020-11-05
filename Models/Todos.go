package Models

import (
	"../Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// fetch all todos at once and
func GetAllTodos(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}

	return nil
}

// insert a todo into the database
func CreateATodo(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}

	return nil
}

// fetch one todo from the database
func GetATodo(todo *Todo, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}

	return nil
}

// update a todo to the database
func UpdateATodo (todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

// delete a todo from the database
func DeleteATodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}