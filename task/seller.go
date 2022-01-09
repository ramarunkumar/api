package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func createseller(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	name := c.PostForm("name")
	email := c.PostForm("email")
	phoneno := c.PostForm("phoneno")

	fmt.Println(name)

	if _, err := sellervalid(name, email, phoneno); err == nil {
		rows, err := db.Query("INSERT INTO seller(name, email,phoneno) VALUES('" + name + "','" + email + "','" + phoneno + "')")
		if rows != nil {
			fmt.Println("error", rows)
		} else {
			fmt.Println("insert error", err)
		}

		for rows.Next() {
			emp := Seller{}

			err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
			if err != nil {
				fmt.Println("scan error", err)
			}

		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"name":    name,
			"email":   email,
			"phoneno": phoneno,
			"role":    2,
		})

	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{

			"Message": err.Error()})

	}
}

func sellervalid(name, email, phoneno string) (*Seller, error) {

	if selleremailavailable(email) {
		return nil, errors.New("email not available")
	}
	if !strings.Contains(email, "@") {

		return nil, errors.New("email must have symbol @")
	}
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("the email can't be empty")
	}
	if strings.TrimSpace(phoneno) == "" {
		return nil, errors.New("the phonenumber can't be empty")
	}

	if phonenoavailableseller(phoneno) {
		return nil, errors.New("phone number already exits")
	}

	if len(phoneno) != 10 {
		return nil, errors.New("phone number  only 10 digit ")

	}
	u := Seller{Name: name, Email: email, Phoneno: phoneno}

	return &u, nil
}

func selleremailavailable(email string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	stmt := "SELECT email FROM seller WHERE email = ('" + email + "')"
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

func phonenoavailableseller(phoneno string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	stmt := "SELECT phoneno FROM seller WHERE phoneno = ('" + phoneno + "')"
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
func getAllSeller(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	rows, err := db.Query("SELECT * FROM seller")
	if err != nil {
		fmt.Println("error")
	}
	res := []Seller{}

	for rows.Next() {
		emp := Seller{}

		err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)
	}
	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, res)
}

func getsellerId(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	id := c.Param("id")
	fmt.Println(id)
	res := []Seller{}
	fmt.Println(res)
	rows, err := db.Query("SELECT * FROM seller")
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

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sellerid not found"})

}
