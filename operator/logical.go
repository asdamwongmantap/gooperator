package operator

import (
	Valuetype "gooperator/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logical(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Arithmethic
	c.BindJSON(&inputvalue)
	var value1 = inputvalue.A
	var value2 = inputvalue.B

	result1 := "Your first value is not good !"
	result2 := "Your second value is not good !"

	if value1 > 4 && value1 < 10 {
		result1 = "Your first value is between 4 and 10 !"
	}

	if value2 == 3 || value2 >= 4 {
		result2 = "Your second value is equal 3 or greater or equal 4 !"
	}

	c.JSON(http.StatusOK, gin.H{"result1": result1, "result2": result2})

}
