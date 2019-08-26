// server2 is a minimal "echo" server and counter server 
package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type Class struct {
	id    int
	name  string
}



func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "gotest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

//func Index(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Class{}
//	res := []Class{}
//	for selDB.Next() {
//		var id int
//		var name string
//		err = selDB.Scan(&id, &name)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//		res = append(res, emp)
//	}
//	tmpl.ExecuteTemplate(w, "Index", res)
//	defer db.Close()
//}
//
//func Show(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Show", emp)
//	defer db.Close()
//}
//
//
//
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}
//
//func Edit(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Edit", emp)
//	defer db.Close()
//}

func Insert(w http.ResponseWriter, r *http.Request) {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/gotest")

	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	class := new(Class)
	class.name = "Trần Bá"

	fmt.Println(o.Insert(class))

	//db := dbConn()
	//if r.Method == "POST" {
	//	name := r.FormValue("name")
	//	var class Class = Class{
	//		name: name,
	//	}
	//	insForm, err := db.Prepare("INSERT INTO Class SET ?")
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	insForm.Exec(class)
	//	log.Println("INSERT: Name: " + name )
	//}
	//defer db.Close()
	//http.Redirect(w, r, "/", 301)
}
//
//func Update(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	if r.Method == "POST" {
//		name := r.FormValue("name")
//		city := r.FormValue("city")
//		id := r.FormValue("uid")
//		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
//		if err != nil {
//			panic(err.Error())
//		}
//		insForm.Exec(name, city, id)
//		log.Println("UPDATE: Name: " + name + " | City: " + city)
//	}
//	defer db.Close()
//	http.Redirect(w, r, "/", 301)
//}
//
//func Delete(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	emp := r.URL.Query().Get("id")
//	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
//	if err != nil {
//		panic(err.Error())
//	}
//	delForm.Exec(emp)
//	log.Println("DELETE")
//	defer db.Close()
//	http.Redirect(w, r, "/", 301)
//}

func main() {
	init()
	log.Println("Server started on: http://localhost:8080/new")
	//http.HandleFunc("/", Index)
	//http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	//http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	//http.HandleFunc("/update", Update)
	//http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}

