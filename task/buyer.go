package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func createbuyer(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	name := c.PostForm("name")
	email := c.PostForm("email")
	phoneno := c.PostForm("phoneno")
	if _, err := buyervalid(name, email, phoneno); err == nil {
		rows, err := db.Query("INSERT INTO buyer (name, email,phoneno) VALUES ('" + name + "','" + email + "','" + phoneno + "')")
		if rows != nil {
			fmt.Println("error", err)
		}
		res := []Buyer{}

		for rows.Next() {
			emp := Buyer{}

			err = rows.Scan(&emp.Id, &emp.Name, &emp.Phoneno, &emp.Role)
			if err != nil {
				fmt.Println("scan error", err)
			}
			fmt.Println(emp.Name, emp.Phoneno, emp.Email)
			res = append(res, emp)
		}
		fmt.Println(res)

		c.IndentedJSON(http.StatusOK, gin.H{
			"name":    name,
			"email":   email,
			"phoneno": phoneno,
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{

			"Message": err.Error()})

	}
}

func buyervalid(name, email, phoneno string) (*Seller, error) {
	if !strings.Contains(email, "@") {

		return nil, errors.New("email must have symbol @")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("the email can't be empty")
	}
	if strings.TrimSpace(phoneno) == "" {
		return nil, errors.New("the phonenumber can't be empty")
	}
	if buyeremailavailable(email) {
		return nil, errors.New("email not available")
	}
	if phonenoavailablebuyer(phoneno) {
		return nil, errors.New("phone number already exits")
	}

	if len(phoneno) != 10 {
		return nil, errors.New("phone number only 10 digit")

	}
	u := Seller{Name: name, Email: email, Phoneno: phoneno}

	return &u, nil
}

func buyeremailavailable(email string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	stmt := "SELECT email FROM buyer WHERE email = ('" + email + "')"
	fmt.Println(stmt)
	err = db.QueryRow(stmt).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("email error", err)
		}
		return false
	}

	return true
}

func phonenoavailablebuyer(phoneno string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	stmt := "SELECT phoneno FROM buyer WHERE phoneno = ('" + phoneno + "')"
	fmt.Println(stmt)
	err = db.QueryRow(stmt).Scan(&phoneno)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("email error", err)
		}
		return false
	}

	return true
}

func getAllBuyers(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	rows, err := db.Query("SELECT * FROM buyer")
	if rows != nil {
		fmt.Println("error", err)
	}
	res := []Buyer{}
	for rows.Next() {
		emp := Buyer{}

		err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)
	}
	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}

func getbuyerId(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	id := c.Param("id")
	fmt.Println(id)
	res := []Buyer{}
	fmt.Println(res)
	rows, err := db.Query("SELECT * FROM buyer")
	if err != nil {
		fmt.Println("error")
	}
	for rows.Next() {
		emp := Seller{}

		err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
		if err != nil {
			fmt.Println("scan error", err)
		}
		if id == emp.Id {
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Buyerid is not found"})

}
