package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Res - Response struct
type Res struct {
	Ctx    *gin.Context `json:"-"`
	Err    error        `json:"error,omitempty"`
	Msg    string       `json:"message,omitempty"`
	Status int          `json:"status,omitempty"`
	Data   interface{}  `json:"data,omitempty"`
}

// ResErr - sends a json response if there is an error
func ResErr(r Res) {
	if r.Msg == "" {
		r.Msg = "Error occured"
	}
	if r.Status == 0 {
		r.Status = 400
	}
	log.Println("error: ", r.Err)
	r.Ctx.JSON(r.Status, r)
}

// ResSuccess - sends a json respnse on success
func ResSuccess(r Res) {
	if r.Msg == "" {
		r.Msg = "Successful!"
	}
	if r.Status == 0 {
		r.Status = 200
	}
	r.Ctx.JSON(r.Status, r)
}
