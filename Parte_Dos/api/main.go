package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	EmployeeNumber int    `json:"employeeNumber"`
	LastName       string `json:"lastName"`
	FirstName      string `json:"firstName"`
	Email          string `json:"email"`
	JobTitle       string `json:"jobTitle"`
}

type Customer struct {
	CustomerNumber   int    `json:"customerNumber"`
	CustomerName     string `json:"customerName"`
	ContactLastName  string `json:"contactLastName"`
	ContactFirstName string `json:"contactFirstName"`
	Phone            string `json:"phone"`
}

var db *sql.DB

func initDB() {
	var err error
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL database")
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT employeeNumber, lastName, firstName, email, jobTitle FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.EmployeeNumber, &e.LastName, &e.FirstName, &e.Email, &e.JobTitle); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT customerNumber, customerName, contactLastName, contactFirstName, phone FROM customers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.CustomerNumber, &c.CustomerName, &c.ContactLastName, &c.ContactFirstName, &c.Phone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		customers = append(customers, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/customers", getCustomers)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
