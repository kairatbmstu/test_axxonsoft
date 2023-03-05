package controller

import (
	"fmt"

	"example.com/test_axxonsoft/v2/dto"
	"github.com/gin-gonic/gin"
)

func PostTask(c *gin.Context) {

}

func GetTask(c *gin.Context) {
	fmt.Println("id : " + c.GetString("id"))
	var taskDto = dto.TaskDTO{
		Id:             c.GetString("id"),
		Method:         "GET",
		HttpStatusCode: "OK",
		Url:            "http://google.com",
		ResponseBody:   "{message:ok}",
	}

	c.JSON(200, taskDto)
}
