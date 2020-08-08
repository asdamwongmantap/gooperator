package operator

import (
	Valuetype "gooperator/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Comparison(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Arithmethic
	c.BindJSON(&inputvalue)
	var value1 = inputvalue.A
	var value2 = inputvalue.B

	result1 := "Your first value is not good !"
	result2 := "Your second value is not good !"

	if value1 > 4 {
		result1 = "Your first value is greater than 4 !"
	} else if value1 >= 5 {
		result1 = "Your first value is greater or equal than 4 !"
	} else if value1 < 10 {
		result1 = "Your first value is less than 10 !"
	} else if value1 <= 8 {
		result1 = "Your first value is less or equal than 8 !"
	}

	if value2 == 3 {
		result2 = "Your second value is equal 3 !"
	} else if value2 != 2 {
		result2 = "Your second value is not equal 2 !"
	}

	c.JSON(http.StatusOK, gin.H{"result1": result1, "result2": result2})

}
