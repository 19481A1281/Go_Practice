package controllers

import (
	"github.com/19481A1281/employee-department/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/19481A1281/employee-department/models"
	// "gorm.io/gorm"
	// "strconv"
)

type DepartmentController struct {
	service *services.DepartmentService
}

func NewDepartmentController(s *services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: s}
}

func (c *DepartmentController) RegisterRoutes(r *gin.Engine) {
	deptGroup := r.Group("/department")
	deptGroup.POST("/", c.CreateDepartment)
	deptGroup.GET("/", c.GetAllDepartments)
}

func (c *DepartmentController) CreateDepartment(ctx *gin.Context) {
	var dept models.Department
	// Bind the incoming JSON to the struct
	if err := ctx.ShouldBindJSON(&dept); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateDepartment(&dept); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dept)
}

func (c *DepartmentController) GetAllDepartments(ctx *gin.Context) {
	departments, err := c.service.GetAllDepartments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, departments)
}