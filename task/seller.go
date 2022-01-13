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
	db := dbinit()

	var emp Users
	if err := c.ShouldBindJSON(&emp); err != nil {
		fmt.Println("error", err)
	}
	Name := emp.Name
	Email := emp.Email
	Phoneno := emp.Phoneno

	fmt.Println("name", Name, Email, Phoneno)
	if _, err := sellervalid(Name, Email, Phoneno); err == nil {
		var res Users
		role := "2"
		rows := db.QueryRow("INSERT INTO users(name, email,phoneno,role) VALUES('"+Name+"','"+Email+"','"+Phoneno+"','"+role+"')").Scan(&res.Id, &res.Name, &res.Email, &res.Phoneno, &res.Role)
		if rows != nil {
			fmt.Println("inserted", rows)
		}
		data := "successfully registered seller account"
		c.IndentedJSON(http.StatusOK, gin.H{
			data: emp.Email,
		})

	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": err.Error()})
	}
}

func sellervalid(name, email, phoneno string) (*Users, error) {

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
	u := Users{Name: name, Email: email, Phoneno: phoneno}

	return &u, nil
}

func selleremailavailable(email string) bool {
	db := dbinit()
	stmt := "SELECT email FROM users WHERE email = ('" + email + "')"
	fmt.Println(stmt)
	err := db.QueryRow(stmt).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("email error", err)
		}
		return false
	}

	return true
}

func phonenoavailableseller(phoneno string) bool {
	db := dbinit()
	stmt := "SELECT phoneno FROM users WHERE phoneno = ('" + phoneno + "')"
	fmt.Println(stmt)
	err := db.QueryRow(stmt).Scan(&phoneno)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("email error", err)
		}
		return false
	}

	return true
}

//-----------------------------------------------------getAllseller------------------------------------//

func getAllSeller(c *gin.Context) {
	db := dbinit()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("error")
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

//-------------------------------------------getsellerId---------------------------------------//

func getsellerId(c *gin.Context) {
	db := dbinit()
	id := c.Param("id")
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

		if id == emp.Id {
			c.IndentedJSON(http.StatusOK, emp)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sellerid not found"})

}
