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

	router.POST("/createproduct/seller:seller_id", createproduct)
	router.GET("/product/:id", getproductID)
	router.GET("/buyproduct", buyproduct)
	router.Run()
}

func homepage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "E-commerce API")
}
