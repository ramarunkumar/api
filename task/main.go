package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.GET("/", homepage)

	router.POST("/seller/create", createseller)
	router.GET("/sellers", getAllSeller)
	router.GET("/seller/:id", getsellerId)

	router.POST("/buyer/create", createbuyer)
	router.GET("/buyers", getAllBuyers)
	router.GET("/buyer/:id", getbuyerId)
	//----------------------------finished--------------------------------------//

	router.POST("/createproduct/seller/:seller", createproduct)
	router.GET("/orderproduct/buyer", orderproduct)
	router.GET("/product/:id", getproductID)

	router.Run()
}

func homepage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "E-commerce API")
}
