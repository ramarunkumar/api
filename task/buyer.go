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

	var emp Users
	if err := c.ShouldBindJSON(&emp); err != nil {
		fmt.Println("error", err)
	}
	name := emp.Name
	email := emp.Email
	phoneno := emp.Phoneno

	if _, err := buyervalid(name, email, phoneno); err == nil {
		role := "1"
		rows, err := db.Query("INSERT INTO users (name, email,phoneno,role) VALUES ('" + name + "','" + email + "','" + phoneno + "','" + role + "')")
		if rows != nil {
			fmt.Println("error", err)
		}
		res := []Users{}

		for rows.Next() {
			emp := Users{}

			err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
			if err != nil {
				fmt.Println("scan error", err)
			}
			fmt.Println(emp.Name, emp.Phoneno, emp.Email)
			res = append(res, emp)
		}
		fmt.Println(res)
		data := "successfully registered buyer account"
		c.IndentedJSON(http.StatusOK, gin.H{
			data: emp.Email,
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{

			"Message": err.Error()})

	}
}

func buyervalid(name, email, phoneno string) (*Users, error) {
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
	u := Users{Name: name, Email: email, Phoneno: phoneno}

	return &u, nil
}

func buyeremailavailable(email string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/api")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	stmt := "SELECT email FROM users WHERE email = ('" + email + "')"
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
	stmt := "SELECT phoneno FROM users WHERE phoneno = ('" + phoneno + "')"
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

	rows, err := db.Query("SELECT * FROM users")
	if rows != nil {
		fmt.Println("error", err)
	}
	res := []Users{}
	for rows.Next() {
		emp := Users{}

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
	id := c.Param("role")
	fmt.Println(id)
	res := []Users{}
	fmt.Println(res)
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("error")
	}
	for rows.Next() {
		emp := Users{}

		err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
		if err != nil {
			fmt.Println("scan error", err)
		}
		if id == emp.Role {
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Buyerid is not found"})

}
