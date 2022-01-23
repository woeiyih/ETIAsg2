package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//For Student's Microservice
var students Students
var StudentID string
var db *sql.DB

type Students struct {
	StudentID   string
	Name        string
	Description string
}

func CreateNewStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("INSERT INTO Students VALUES ('%s', '%s', '%s')",
		s.StudentID, s.Name, s.Description)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func UpdateStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("UPDATE Students SET Name='%s', Description='%s' WHERE StudentID='%s'",
		s.Name, s.Description, s.StudentID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func ViewStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE StudentID='%s'",
		s.StudentID, s.Name, s.Description)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("DELETE FROM Students WHERE StudentID='%s'",
		s.StudentID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func ListStudents(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students",
		s.StudentID, s.Name, s.Description)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func SearchStudents(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE Name='%s'",
		s.StudentID, s.Name, s.Description)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func student(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-type") == "application/json" {
		// POST is for creating new passenger
		if r.Method == "POST" {
			// read the string sent to the service
			var newStudent Students
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newStudent)
				//Check if user fill up the required information for registering Passenger's account
				if newStudent.StudentID == "" || newStudent.Name == "" || newStudent.Description == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply passenger " + "information " + "in JSON format"))
					return
				} else {
					CreateNewStudent(db, newStudent) //Once everything is checked, passenger's account will be created
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Successfully created passenger's account"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information " +
					"in JSON format"))
			}
		}
		//---PUT is for creating or updating existing passenger---
		if r.Method == "PUT" {
			queryParams := r.URL.Query() //used to resolve the conflict of calling API using the '%s'?PassengerID='%s' method
			StudentID = queryParams["StudentID"][0]
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &students)
				//Check if user fill up the required information for updating Passenger's account information
				if students.StudentID == "" || students.Name == "" || students.Description == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply passenger " + " information " + "in JSON format"))
				} else {
					students.StudentID = StudentID
					UpdateStudent(db, students)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Successfully updated passenger's information"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "passenger information " + "in JSON format"))
			}
		}

	}
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
	//---Deny any deletion of passenger's account or other passenger's information
	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - For audit purposes, passenger's account cannot be deleted."))
	}
}

func liststudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
}

func searchstudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
}

func main() {
	// instantiate passengers
	ridesharing_db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ridesharing_db")

	db = ridesharing_db
	// handle error
	if err != nil {
		panic(err.Error())
	}
	//handle the API connection across all three microservices, Passengers, Trips and Drivers
	router := mux.NewRouter()
	router.HandleFunc("/students", student).Methods(
		"GET", "POST", "PUT", "DELETE")
	router.HandleFunc("/students/liststudents", liststudent).Methods(
		"GET")
	router.HandleFunc("/students/searchstudents", searchstudent).Methods(
		"GET")
	fmt.Println("Students microservice API --> Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))

	defer db.Close()
}
