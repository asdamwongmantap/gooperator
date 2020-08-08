package main

import (
	"log"
	"net/http"

	Conf "gooperator/config"
	Coperator "gooperator/operator"

	"github.com/gin-gonic/gin"

	"github.com/rs/cors"
)

func main() {

	addr, err := Conf.MyPort()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/gooperator")
	{
		v1.POST("/arithmeticopr", Coperator.Arithmethic)
		v1.POST("/comparisonopr", Coperator.Comparison)
		v1.POST("/logicalopr", Coperator.Logical)
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}
