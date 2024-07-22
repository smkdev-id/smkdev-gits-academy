package route

import (
	"github.com/gin-gonic/gin"
	"github.com/HerlyRyan/be-task-2a/controller"
)

type App struct {
	Router *gin.Engine
	TodoController controller.TodoController
}

func (a *App) InitializeRoutes() {
	a.Router = gin.Default()
	a.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello from route /"})
	})
	a.Router.POST("/todos", a.TodoController.CreateTodo)
	a.Router.GET("/todos", a.TodoController.GetAllTodos)
	a.Router.GET("/todos/:id", a.TodoController.GetTodoByID)
	a.Router.PUT("/todos/:id", a.TodoController.UpdateTodo)
	a.Router.DELETE("/todos/:id", a.TodoController.DeleteTodo)
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
