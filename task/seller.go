package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"strings"

	"github.com/gin-gonic/gin"
)

//-----------------------------------------------createseller--------------------------------------------------//

func createseller(c *gin.Context) {
	db := dbinit()
	var emp Users
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}

	if _, err := sellervalid(emp.Name, emp.Email, emp.Phoneno); err == nil {
		role := "2"
		rows, err := db.Query("INSERT INTO users (name, email,phoneno,role) VALUES ('" + emp.Name + "','" + emp.Email + "','" + emp.Phoneno + "','" + role + "')RETURNING id,name,email,phoneno,role")
		if rows != nil {
			fmt.Println("error", err)
		}
		fmt.Println(emp.Id)
		res := []Users{}
		fmt.Println(emp.Id)
		for rows.Next() {
			emp := Users{}

			err = rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Phoneno, &emp.Role)
			if err != nil {
				fmt.Println("scan error", err)
			}
			fmt.Println(emp.Name, emp.Phoneno, emp.Email)
			res = append(res, emp)
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data":    res,
			"Message": "successfully registered seller account",
		})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{

			"Message": err.Error()})

	}
}

//----------------------------------------sellervalid--------------------------------------------------//

func sellervalid(name, email, phoneno string) (*Users, error) {

	if selleremailavailable(email) {
		return nil, errors.New("email not available")
	}
	if !isvalid(email) {
		return nil, errors.New("invalid email type")
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
func isvalid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

//------------------------------------------selleremailavailable---------------------------------------//

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

//------------------------------------------phonenoavailableseller---------------------------------------//

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
	rows, err := db.Query("SELECT * FROM users where role=2")
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
	rows, err := db.Query("SELECT * FROM users where role=2 and id='" + id + "'")
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
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "sellerid not found"})
}
