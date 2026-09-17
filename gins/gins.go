package gins

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type Code int

func response(c *gin.Context, r Response) {
	c.JSON(200, r)
}
