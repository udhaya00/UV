
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type ContactDetails1 struct {
	name   string
}


func main(){
	var (
		name string
		dob string
		pre string
	)
	db, err := sql.Open("mysql", "root:kavin2000@(localhost:3306)/employee")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	tmp2 := template.Must(template.ParseFiles("forms1.html"))



	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmp2.Execute(w, nil)
			return
		}

		details1 := ContactDetails1{
			name:   r.FormValue("email"),
		}
		query := "SELECT Name, dob, present FROM birth WHERE Name = ?"
		if err := db.QueryRow(query, details1.name).Scan(&name, &dob, &pre); err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n\tBirthday\t")
		fmt.Printf("..............................")
		fmt.Println("\nName:",name)
		fmt.Printf("..............................")
		fmt.Println("\nDOB:",dob)
		fmt.Printf("..............................")
		fmt.Println("\nGift:",pre)
		fmt.Printf("..............................")
		tmp2.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}