package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/melardev/GoGinGonicApiGormCrud/dtos"
	"github.com/melardev/GoGinGonicApiGormCrud/services"
	"net/http"
	"strconv"
)

func GetAllTodos(c *gin.Context) {
	todos := services.FetchTodos()

	c.JSON(http.StatusOK, dtos.GetTodoListDto(todos))
}

func GetAllPendingTodos(c *gin.Context) {
	todos := services.FetchPendingTodos()
	c.JSON(http.StatusOK, dtos.GetTodoListDto(todos))
}
func GetAllCompletedTodos(c *gin.Context) {
	todos := services.FetchCompletedTodos()
	c.JSON(http.StatusOK, dtos.GetTodoListDto(todos))
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	if id == "completed" {
		GetAllCompletedTodos(c)
		return
	} else if id == "pending" {
		GetAllPendingTodos(c)
		return
	}
	id64, _ := strconv.ParseUint(id, 10, 32)
	todo, err := services.FetchById(uint(id64))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not find Todo"))
		return
	}

	c.JSON(http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func CreateTodo(c *gin.Context) {
	var json dtos.CreateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}
	todo, err := services.CreateTodo(json.Title, json.Description, json.Completed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.GetTodoDetaislDto(&todo))
}

func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}

	var json dtos.CreateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, dtos.CreateBadRequestErrorDto(err))
		return
	}

	todo, err := services.UpdateTodo(uint(id), json.Title, json.Description, json.Completed)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.GetTodoDetaislDto(&todo))

}

func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}
	todo, err := services.FetchById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func DeleteAllTodos(c *gin.Context) {
	services.DeleteAllTodos()
	c.JSON(http.StatusNoContent, nil)
}
