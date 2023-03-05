package controller

import (
	"fmt"
	"net/http"

	"example.com/test_axxonsoft/v2/service"
	"github.com/gin-gonic/gin"
)

func PostTask(c *gin.Context) {

}

func GetTask(c *gin.Context) {
	var id = c.GetString("id")
	fmt.Println("id : " + id)

	task, err := service.TaskServiceInst.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":        "internal_server_error",
			"errorMessage": err.Error(),
		})
	}

	c.JSON(200, task)
}
