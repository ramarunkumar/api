package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dbinit()
	router := gin.Default()

	router.GET("/", homepage)

	router.POST("/createseller", createseller)
	router.GET("/sellers", getAllSeller)
	router.GET("/seller/:id", getsellerId)

	router.POST("/createbuyer", createbuyer)
	router.GET("/buyers", getAllBuyers)
	router.GET("/buyer/:id", getbuyerId)

	router.POST("/createproduct", createproduct)
	router.POST("/orderproduct", orderproduct)
	router.GET("/products", getAllproducts)
	router.GET("/product/:id", getproductID)
	router.Run()
}

func homepage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "E-commerce API")
}

//---------------------------------------------dbconnection-----------------------------//

func dbinit() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	return db
}
