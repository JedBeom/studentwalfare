package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

//Student contains information of a student.
type Student struct {
	ID           int `json:"id"`
	card         string
	Number       int    `json:"number"`
	Name         string `json:"name"`
	Points       int    `json:"points"`
	IsAdmin      bool
	Email        string
	Password     string
	Coupons      int
	ErrorMessage string
}

//History contains what happens in the server.
type History struct {
	Sequence    int
	ID          int
	Location    string
	Description string
	AdminID     int
	ScanTime    string
}

func (history *History) Create() (err error) {
	statement := "insert into history (id, location, description, adminid, scantime) values ($1, $2, $3, $4, $5) returning sequence "
	stmt, err := Db.Prepare(statement)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(history.ID, history.Location, history.Description, history.AdminID, history.ScanTime).Scan(&history.Sequence)
	if err != nil {
		log.Println(err)
	}
	return
}

//데이터베이스를 위한 변수
var Db *sql.DB

//프로그램이 실행될 때 맨 처음에 실행되는 함수로, 데이터베이스 연결을 함.
func init() {
	var err error
	Db, err = sql.Open("postgres", "user="+Username()+" dbname=students password="+Password()+" sslmode=disable")
	//프로그램 실행 전 같은 위치의 const.go 파일 내의 문자열을 수정해야함.
	if err != nil {
		panic(err)
		//여기서 에러나면 프로그램 구동 실패.
	}
}

func getAStudent(id int) (student Student, err error) {
	student = Student{}
	err = Db.QueryRow("select id, number, name, points, is_admin, email, password, coupons from student where id = $1", id).Scan(&student.ID, &student.Number,
		&student.Name, &student.Points, &student.IsAdmin, &student.Email, &student.Password, &student.Coupons)
	return
}

func addPoints(id int, addedInt int) (err error) {
	student := Student{}
	student.ID = id
	err = Db.QueryRow("select points from student where id = $1", id).Scan(&student.Points)
	log.Println(student.Points)
	student.Points = addedInt + student.Points
	log.Println(student.Points)
	_, err = Db.Exec("update student set points = $2 where id = $1", student.ID, student.Points)
	return
}

func SearchByID(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/findbyid.html")
	strid := r.FormValue("id")
	var student Student
	var err2 error
	if strid != "" {
		id, err := strconv.Atoi(strid)
		student, err2 = getAStudent(id)
		if err2 == sql.ErrNoRows {
			student.ErrorMessage = strid + "는(은) 존재하지 않습니다."
		}
		if err == strconv.ErrSyntax {
			student.ErrorMessage = "오류!"
		}
	}
	t.Execute(w, student)
}

func main() {
	/*
		server := http.Server{
			Addr: "127.0.0.1:8080",
		}
		http.HandleFunc("/", SearchByID)
		server.ListenAndServe()
	*/
	t := time.Now()
	timenow := fmt.Sprintf("%v-%v-%v %v:%v:%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println(timenow)
	history := History{ID: 1, Location: "Gym", Description: "Entering", AdminID: 1, ScanTime: timenow}
	fmt.Println(history)
	history.Create()
	fmt.Println(history)
}

/***
 *           _          _ ____
 *          | |        | |  _ \
 *          | | ___  __| | |_) | ___  ___  _ __ ___
 *      _   | |/ _ \/ _` |  _ < / _ \/ _ \| '_ ` _ \
 *     | |__| |  __/ (_| | |_) |  __/ (_) | | | | | |
 *      \____/ \___|\__,_|____/ \___|\___/|_| |_| |_|
 *
 *
 */
