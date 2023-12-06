package main

import (
	"CRUD-API-GO/controllers"
	"CRUD-API-GO/initializers"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVars()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.POST("/employee", controllers.CreateEmployee)
	r.GET("/employee", controllers.GetEmployee)
	r.GET("/employee/:id", controllers.GetEmployeeById)
	r.PUT("/employee/:id", controllers.UpDateEmployee)
	r.DELETE("/employee/:id", controllers.DeleteEmployee)
	r.Run()
}
