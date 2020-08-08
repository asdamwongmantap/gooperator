package operator

import (
	Valuetype "gooperator/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Arithmethic(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Arithmethic
	c.BindJSON(&inputvalue)
	var value1 = inputvalue.A
	var value2 = inputvalue.B

	var addition = value1 + value2
	var subtraction = value1 - value2
	var multiplication = value1 * value2
	var division = value1 / value2
	var modulus = value1 % value2

	c.JSON(http.StatusOK, gin.H{"addition": addition, "subtraction": subtraction,"multiplication":multiplication,
    "division":division,"modulus":modulus})

}
