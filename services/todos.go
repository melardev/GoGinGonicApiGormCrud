package services

import (
	"github.com/melardev/GoGinGonicApiGormCrud/infrastructure"
	"github.com/melardev/GoGinGonicApiGormCrud/models"
)

func FetchTodos() []models.Todo {
	var todos []models.Todo
	database := infrastructure.GetDb()
	database.Select("id, title, completed, created_at, updated_at").
		Find(&todos)
	return todos
}

func FetchPendingTodos() (todos []models.Todo) {
	return FetchTodosByCompleted(false)
}

func FetchCompletedTodos() (todos []models.Todo) {
	return FetchTodosByCompleted(true)
}

func FetchTodosByCompleted(completed bool) (todos []models.Todo) {
	database := infrastructure.GetDb()
	database.Select("id, title, completed, created_at, updated_at").
		// Why this would not work http://doc.gorm.io/crud.html#query
		// Where(&models.Todo{Completed: completed}).
		// whereas the below does work ;)
		Where("completed = ?", completed).
		Order("created_at desc").
		Find(&todos)

	return todos
}

func DeleteAllTodos() {
	database := infrastructure.GetDb()
	database.Model(&models.Todo{}).Delete(&models.Todo{})
}

func FetchById(id uint) (todo models.Todo, err error) {
	database := infrastructure.GetDb()
	err = database.Model(&models.Todo{}).First(&todo, id).Error
	return
}

func DeleteTodo(todo *models.Todo) error {
	database := infrastructure.GetDb()
	return database.Delete(todo).Error
}

func CreateTodo(title, description string, completed bool) (todo models.Todo, err error) {
	database := infrastructure.GetDb()
	todo = models.Todo{Title: title, Description: description, Completed: completed}
	err = database.Create(&todo).Error
	return todo, err
}

func UpdateTodo(id uint, title, description string, completed bool) (todo models.Todo, err error) {
	todo, err = FetchById(id)
	if err != nil {
		return
	}
	todo.Title = title

	// TODO: handle this in a better way, the user should be able to set description to empty string
	// The intention is to check against nil but in go there are no nil strings, so we can not know
	// if the user intended to udpate the description to empty string or just update the other fields other than description.
	if description != "" {
		todo.Description = description
	}

	todo.Completed = completed
	database := infrastructure.GetDb()

	// For updating you can:
	// database.Model(&todo).Updates(map[string]interface{}{"title": title, "description": description, "completed": completed})
	// database.Model(&todo).Updates(models.Todo{Title: title, Description: description, Completed: completed})

	// But I stick with the versatile Save function
	database.Save(&todo)

	return
}
