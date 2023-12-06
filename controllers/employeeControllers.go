package controllers

import (
	"CRUD-API-GO/initializers"
	"CRUD-API-GO/models"
	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	///req body data
	var body struct {
		Name           string
		SecondName     string
		Age            int
		WorkExperience string
	}

	c.Bind(&body)
	employee := models.Employee{
		Name: body.Name, SecondName: body.SecondName, Age: body.Age, WorkExperience: body.WorkExperience,
	}

	result := initializers.DB.Create(&employee)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"employee": employee,
	})
}

func GetEmployee(c *gin.Context) {
	var employee []models.Employee
	//get all
	initializers.DB.Find(&employee)

	c.JSON(200, gin.H{
		"employee": employee,
	})
}
func GetEmployeeById(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	//find by id
	initializers.DB.First(&employee, id)

	c.JSON(200, gin.H{
		"employee": employee,
	})
}
func UpDateEmployee(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name           string
		SecondName     string
		Age            int
		WorkExperience string
	}
	c.Bind(&body)

	var employee models.Employee
	initializers.DB.First(&employee, id)
	initializers.DB.Model(&employee).Updates(models.Employee{
		Name:           body.Name,
		SecondName:     body.SecondName,
		Age:            body.Age,
		WorkExperience: body.WorkExperience,
	})

	c.JSON(200, gin.H{
		"employee": employee,
	})
}
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Employee{}, id)

	c.Status(200)
}
