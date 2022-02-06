package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//For Student's Microservice
//var students map[string]Students
var students Students
var StudentID string
var db *sql.DB

type Students struct {
	StudentID   string
	StudentName string
	DOB         string
	Address     string
	PhoneNumber string
}

func CreateNewStudent(db *sql.DB, s Students) {

	query := fmt.Sprintf("INSERT INTO Students VALUES ('%s', '%s', '%s', '%s', '%s')",
		s.StudentID, s.StudentName, s.DOB, s.Address, s.PhoneNumber)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func UpdateStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("UPDATE Students SET StudentName='%s', DOB='%s', Address='%s', s.PhoneNumber='%s' WHERE StudentID='%s'",
		s.StudentName, s.DOB, s.Address, s.PhoneNumber, s.StudentID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func ViewStudent(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE StudentID='%s'",
		s.StudentID)
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
func ListStudents(db *sql.DB) []Students {
	/*if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}*/
	results, err := db.Query("SELECT StudentID, StudentName FROM Students")

	if err != nil {
		panic(err.Error())
	}
	var students []Students
	for results.Next() {
		var getStudent Students
		err = results.Scan(&getStudent.StudentID, &getStudent.StudentName)
		if err != nil {

			panic(err.Error())
		}
		students = append(students, getStudent)
	}
	return students

}

/* func ListStudents(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students",
		s.StudentID, s.StudentName, s.DOB, s.Address, s.PhoneNumber)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
*/

func SearchStudents(db *sql.DB, s Students) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE Name='%s'",
		s.StudentName)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

/*func searchstudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
}*/

func student(w http.ResponseWriter, r *http.Request) {
	////params := mux.Vars(r)
	if r.Header.Get("Content-type") == "application/json" {
		// POST is for creating new student
		if r.Method == "POST" {
			// read the string sent to the service
			var newStudent Students
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &newStudent)
				//Check user has filled up the required information for registering
				if newStudent.StudentID == "" || newStudent.StudentName == "" || newStudent.DOB == "" || newStudent.Address == "" || newStudent.PhoneNumber == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply student " + "information " + "in JSON format"))
					return
				} else {
					CreateNewStudent(db, newStudent) //Validation before creating
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Successfully created student account"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply student information " +
					"in JSON format"))
			}
		}
		//PUT - create/update existing student
		if r.Method == "PUT" {
			queryParams := r.URL.Query()
			StudentID = queryParams["StudentID"][0]
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &students)
				//Check if user fill up the required information for updating student account information
				if students.StudentID == "" || students.StudentName == "" || students.DOB == "" || students.Address == "" || students.PhoneNumber == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply `student` " + " information " + "in JSON format"))
				} else {
					students.StudentID = StudentID
					UpdateStudent(db, students)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Successfully updated student information"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "student information " + "in JSON format"))
			}
		}

	}
	if r.Method == "GET" {
		/*if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}*/
		StudentID := r.URL.Query().Get("StudentID")
		fmt.Println("StudentID: ", StudentID)
		students := ListStudents(db)

		json.NewEncoder(w).Encode(&students)

	}
	//---Deny any deletion of student's account or other student's information
	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403 - For audit purposes, student account cannot be deleted."))
	}
}

func main() {
	// instantiate students
	a2_db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/a2_db")

	db = a2_db
	// handle error
	if err != nil {
		panic(err.Error())
	}
	//handle the API connection
	router := mux.NewRouter()

	router.HandleFunc("/students", student).Methods("GET", "POST", "PUT", "DELETE")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	fmt.Println("Student Microservice API --> Listening at port 8150")
	log.Fatal(http.ListenAndServe(":8150", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

	/*router.HandleFunc("/students", student).Methods(
		"GET", "POST", "PUT", "DELETE")
	router.HandleFunc("/students/liststudents", ListStudents()).Methods(
		"GET")
	router.HandleFunc("/students/searchstudents", ListStudents()).Methods(
		"GET")
	fmt.Println("Students microservice API --> Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))*/

	defer db.Close()
}
