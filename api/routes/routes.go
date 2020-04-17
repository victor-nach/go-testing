package routes

import (
	"github.com/gin-gonic/gin"
	c "github.com/victor-nach/user-management-go/api/controllers"
	m "github.com/victor-nach/user-management-go/api/middlewares"
)

// Router - returns a gin router
func Router() *gin.Engine {
	router := gin.New()
	router.GET("/users", c.GetAllUsers)
	router.GET("/users/:id", m.CheckUser, c.GetSingleUser)
	router.POST("/users", c.CreateSingleUser)
	router.PATCH("/users/:id", m.CheckUser, c.UpdateSingleUser)
	router.DELETE("/users/:id", m.CheckUser, c.DeleteSingleUser)
	return router
}
