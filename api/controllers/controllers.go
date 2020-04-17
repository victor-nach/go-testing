package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/victor-nach/user-management-go/api/models"
	u "github.com/victor-nach/user-management-go/api/utils"
)

// CreateSingleUser - controller for creating a single user
func CreateSingleUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Invalid request", Err: err})
		return
	}
	if err := user.CreateUser(); err != nil {
		u.ResErr(u.Res{Ctx: c})
		return
	}
	u.ResSuccess(u.Res{Ctx: c, Data: user})
}

// GetSingleUser - controller for returning a single user
func GetSingleUser(c *gin.Context) {
	var user models.User
	user = c.MustGet("user").(models.User)
	u.ResSuccess(u.Res{Ctx: c, Data: user})
}

// GetAllUsers - controller for returning all users
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		u.ResErr(u.Res{Ctx: c})
		return
	}
	u.ResSuccess(u.Res{Ctx: c, Data: users})
}
// UpdateSingleUser - controller for updating a single user
func UpdateSingleUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		u.ResErr(u.Res{Ctx: c, Msg: "Invalid request", Err: err})
		return
	}
	if err := user.UpdateUser(c.Param("id")); err != nil {
		u.ResErr(u.Res{Ctx: c, Err: err})
		return
	}
	u.ResSuccess(u.Res{Ctx: c, Data: user})
}

// DeleteSingleUser - controller for deleting a single user
func DeleteSingleUser(c *gin.Context) {
	var user models.User
	if err := user.DeleteUser(c.Param("id")); err != nil {
		u.ResErr(u.Res{Ctx: c, Err: err})
		return
	}
	u.ResSuccess(u.Res{Ctx: c})
}
