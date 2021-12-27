package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-todo-app/models"
)

func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "status": "notfound"})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "status": "notfound"})
	} else {
		c.BindJSON(&todo)
		err = models.UpdateATodo(&todo)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusAccepted, todo)
		}
	}
}

func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"id": id, "status": "notfound"})
	} else {
		c.BindJSON(&todo)
		err = models.DeleteATodo(&todo)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"id": id, "status": "Deleted"})
		}
	}
}
