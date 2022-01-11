package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	pg := dbconc()
	fmt.Println(pg)
	router.GET("/", homepage)

	router.POST("/seller/create", createseller)
	router.GET("/sellers", getAllSeller)
	router.GET("/seller/:id", getsellerId)

	router.POST("/buyer/create", createbuyer)
	router.GET("/buyers", getAllBuyers)
	router.GET("/buyer/:id", getbuyerId)

	router.POST("/createproduct/seller/:seller", createproduct)
	router.GET("/orderproduct/buyer", orderproduct)
	router.GET("/products", getAllproducts)
	router.GET("/product/:id", getproductID)
	router.Run()
}

func homepage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "E-commerce API")
}

func dbconc() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "qwerty123",
		Addr:     "localhost:5432",
		Database: "api",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		fmt.Println("failed")
		os.Exit(100)
	}
	fmt.Println("successfully")
	closeErr := db.Close()
	if closeErr != nil {
		fmt.Println("error")
		os.Exit(100)
	}
	log.Printf("connected")

	return db
}
