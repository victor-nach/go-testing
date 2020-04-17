package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/victor-nach/user-management-go/api/models"
	u "github.com/victor-nach/user-management-go/api/utils"
)

// CheckUser - middleware for checking if user exists
func CheckUser(c *gin.Context) {
	var user models.User
	if err := user.GetUserByID(c.Param("id")); err != nil {
		u.ResErr(u.Res{Ctx: c, Err: err})
		return
	}
	c.Set("user", user)
}
