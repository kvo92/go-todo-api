package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-todo-app/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// group accessible on /v1
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
	}
	return r
}
