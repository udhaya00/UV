package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Name   string
	dob string
	gift string
}

func main(){

	db, err := sql.Open("mysql", "root:kavin2000@(localhost:3306)/employee")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:   r.FormValue("email"),
			dob: r.FormValue("subject"),
			gift: r.FormValue("message"),
		}
		insert, _ := db.Query("INSERT INTO birth VALUES (?,?,?)",details.Name,details.dob,details.gift)
		defer insert.Close()
		tmpl.Execute(w, struct{ Success bool }{true})
	})


	http.ListenAndServe(":8089", nil)
}