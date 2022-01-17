package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//--------------------------------------------------getProductId-------------------------------------//

func getproductID(c *gin.Context) {
	db := dbinit()
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
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
}

//-------------------------------getAllproducts-------------------------------------------------------//

func getAllproducts(c *gin.Context) {
	db := dbinit()
	rows, err := db.Query("SELECT * FROM products order by id ASC ")
	if err != nil {
		fmt.Println("error")
	}
	res := []Product{}
	for rows.Next() {
		emp := Product{}
		err = rows.Scan(&emp.Id, &emp.Name, &emp.Price, &emp.Tax, &emp.Seller_id)
		if err != nil {
			fmt.Println("scan error", err)
		}
		res = append(res, emp)
	}
	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}

//-----------------------------------------------seller create product----------------------------------//

func createproduct(c *gin.Context) {
	db := dbinit()
	var err error
	var res Product
	if err := c.ShouldBindJSON(&res); err != nil {
		fmt.Println("bind json error", err)
	}
	seller_id := res.Seller_id
	name := res.Name
	price := res.Price
	tax := res.Tax
	fmt.Println(seller_id, name, price, tax)
	fmt.Println("emp", res)
	var sel Users
	id := res.Seller_id
	fmt.Println("id", id)
	err = db.QueryRow("SELECT * FROM users WHERE  users.id='"+id+"'").Scan(&sel.Id, &sel.Name, &sel.Email, &sel.Phoneno, &sel.Role)
	fmt.Println(err)
	fmt.Println("role", sel.Role)
	fmt.Println(sel.Id)
	switch {
	case sel.Role == "2":
		fmt.Println(sel.Role)
		seller_id := id
		fmt.Println("INSERT INTO products( name , price, tax,seller_id)VALUES ('" + name + "', '" + price + "','" + tax + "','" + seller_id + "')")
		rows, err := db.Query("INSERT INTO products( name , price, tax,seller_id)VALUES ('" + name + "', '" + price + "','" + tax + "','" + seller_id + "')RETURNING id,name,price,tax,seller_id")
		if err != nil {
			fmt.Println("insert ", err)
		} else {
			fmt.Println("inserted", rows)
		}
		p := []Product{}
		for rows.Next() {
			emp := Product{}
			err = rows.Scan(&emp.Id, &emp.Name, &emp.Price, &emp.Tax, &emp.Seller_id)
			if err != nil {
				fmt.Println("scan error", err)
			}
			p = append(p, emp)
		}
		data := "successfully added products"
		c.IndentedJSON(http.StatusOK, gin.H{
			data:   name,
			"data": p,
		})

		return

	case err != nil:
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"Message": "seller  id not regestered",
		})
		return
	default:
		c.IndentedJSON(http.StatusOK, "seller  id not regestered")
	}
}

//----------------------------------------------orderproduct-----------------------------------------//

func orderproduct(c *gin.Context) {
	db := dbinit()
	var emp Users
	var res Order
	if err := c.ShouldBindJSON(&res); err != nil {
		fmt.Println("error", err)
	}
	email := res.Email
	name := res.Name
	quantity := res.Quantity
	fmt.Println(res.Name)
	err := db.QueryRow("SELECT * from users where email='"+email+"'").Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
	fmt.Println(res.Price)
	fmt.Println(emp.Role)
	switch {
	case emp.Role == "1":
		fmt.Println(emp.Role)
		err := db.QueryRow("SELECT * from users, products Where users.email='"+email+"'AND products.name='"+name+"'").Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role, &res.Id, &res.Name, &res.Price, &res.Tax, &res.Seller_id)
		if err != nil {
			fmt.Println("no error", err)
		}
		fmt.Println(res.Tax, res.Price, res.Quantity)
		tax, err := strconv.ParseFloat(res.Tax, 64)
		if err != nil {
			fmt.Println(err)
		}
		quan, err := strconv.ParseFloat(quantity, 64)
		if err != nil {
			fmt.Println(err)
		}
		pri, err := strconv.ParseFloat(res.Price, 64)
		if err != nil {
			fmt.Println(err)
		}
		total_price := quan * pri //2*20=40
		fmt.Println("total price", total_price)
		total_tax := (quan * tax) //2*4=8
		fmt.Println("total tax", total_tax)
		// total := total_price + total_tax //40+8=48
		total := total_price * (1 + ((tax) / 100))
		fmt.Println("total", total)

		Message := "You order successfully created"
		c.IndentedJSON(http.StatusOK, gin.H{
			"Name":         res.Name,
			"Tax":          total_tax,
			"Price":        total_price,
			"total amount": total,
			Message:        emp.Email,
		})
		return
	case err != nil:
		c.IndentedJSON(http.StatusOK, gin.H{
			"Message": "buyer email not regester ",
		})
		return
	default:
		c.IndentedJSON(http.StatusOK, "empty")
	}
}
