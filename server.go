package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Student struct {
	Id     int    `json:"id"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=Junhwan dbname=students password=bjhbeom sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func get_a_student(id int) (student Student, err error) {
	student = Student{}
	err = Db.QueryRow("select id, number, name, points from student where id = $1", id).Scan(&student.Id, &student.Number,
		&student.Name, &student.Points)
	return
}

func edit_points(id int, added_int int) (err error) {
	student := Student{}
	student.Id = id
	err = Db.QueryRow("select points from student where id = $1", id).Scan(&student.Points)
	log.Println(student.Points)
	student.Points = added_int + student.Points
	log.Println(student.Points)
	_, err = Db.Exec("update student set points = $2 where id = $1", student.Id, student.Points)
	return
}

func get_handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/student/", get_handler)
	id := 1
	student, err := get_a_student(id)
	if err != nil {
		panic(err)
	}
	output, err2 := json.MarshalIndent(&student, "", "\t\t")
	if err2 != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(output)
}
