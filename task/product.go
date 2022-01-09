package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getproductID(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	id := c.Param("id")
	fmt.Println(id)
	res := []Product{}
	fmt.Println(res)
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println("error")
	}
	for rows.Next() {
		emp := Product{}

		err = rows.Scan(&emp.Id, &emp.Name, &emp.Price, &emp.Tax, &emp.Seller_id)
		if err != nil {
			fmt.Println("scan error", err)
		}
		if id == emp.Id {
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

//-----------------------------------------------seller create product----------------------------------//

func createproduct(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	seller_id := c.Param("seller_id")
	email := c.PostForm("email")
	name := c.PostForm("name")
	price := c.PostForm("price")
	tax := c.PostForm("tax")

	fmt.Println(email)

	rows, err := db.Query("SELECT email FROM seller WHERE email='" + email + "'")
	if rows != nil {
		fmt.Println("selected successfully", rows)
	} else {
		fmt.Println("error", err)
	}
	for rows.Next() {
		emp := Seller{}

		err = rows.Scan(&emp.Email)
		if err != nil {
			fmt.Println("scan error", err)
		}
	}
	emp := Seller{}
	fmt.Println(emp.Email, email)
	if email == emp.Email {
		fmt.Println(emp.Email)
		// select * from seller,products  where seller.id=products.seller_id;
		// INSERT INTO products(name, price, tax, seller_id)VALUES('phone',10000,5,1)
		rows, err = db.Query("INSERT INTO products( name , price, tax,seller_id)	VALUES ('" + name + "', '" + price + "','" + tax + "','" + seller_id + "') ")
		if rows != nil {
			fmt.Println("inserted successfully", err)
		} else {
			fmt.Println("error", err)
		}

		for rows.Next() {
			emp := Product{}

			err = rows.Scan(&emp.Name, &emp.Tax, &emp.Price)
			if err != nil {
				fmt.Println("scan error", err)
			}
			c.IndentedJSON(http.StatusOK, gin.H{
				"name":  name,
				"price": price,
				"tax":   tax,
			})
		}
	} else {

		c.IndentedJSON(http.StatusOK, gin.H{
			"Message": "Email id not regestered",
		})
	}

}

//----------------------------------------------orderproduct-----------------------------------------//

func buyproduct(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	email := c.PostForm("email")
	// quantity := c.PostForm("quantity")
	// name := c.PostForm("name")
	fmt.Println(email)
	rows, err := db.Query("SELECT product_id,  seller_id FROM products JOIN seller_id=1")
	if rows != nil {
		fmt.Println("selected successfully", err)
	} else {
		fmt.Println("error", err)

	}

	for rows.Next() {
		emp := Buyer{}

		err = rows.Scan(&emp.Email)
		if err != nil {
			fmt.Println("scan error", err)
		}
		if email == emp.Email {
			fmt.Println(emp.Email, email)

			c.IndentedJSON(http.StatusOK, gin.H{
				"Message": "You can order your product ",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"Message": "email not regester ",
			})
		}

	}

}
