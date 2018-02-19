package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//Student contains information of a student.
type Student struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

//Db is for a database. See what happens in the main function.
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=username dbname=students password=yourpassword sslmode=disable")
	// You should replace your username and yourpassword.
	if err != nil {
		panic(err)
	}
}

func getAStudent(id int) (student Student, err error) {
	student = Student{}
	err = Db.QueryRow("select id, number, name, points from student where id = $1", id).Scan(&student.ID, &student.Number,
		&student.Name, &student.Points)
	return
}

func editPoints(id int, addedInt int) (err error) {
	student := Student{}
	student.ID = id
	err = Db.QueryRow("select points from student where id = $1", id).Scan(&student.Points)
	log.Println(student.Points)
	student.Points = addedInt + student.Points
	log.Println(student.Points)
	_, err = Db.Exec("update student set points = $2 where id = $1", student.ID, student.Points)
	return
}

//LogURL shows your url use log.Println.
func LogURL(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/student/", LogURL)
	id := 1
	student, err := getAStudent(id)
	if err != nil {
		panic(err)
	}
	output, err2 := json.MarshalIndent(&student, "", "\t\t")
	if err2 != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(output)
	server.ListenAndServe()
}
