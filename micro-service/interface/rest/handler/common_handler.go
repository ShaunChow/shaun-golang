package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Response Common
type Response struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

// OkResponse 200 
func OkResponse(c *gin.Context, code int, message string, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Code: code,
        Message: message,
        Data: data,
    })
}