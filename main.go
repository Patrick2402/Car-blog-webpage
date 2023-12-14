package main

import (
	"fmt"
	"golang/handlers"
	"log"
	"net/http"

	// "fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func routerStart() {

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/blog", handlers.BlogHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("We'h got some problem with running server bro ", err)
	}
}

func databaseConnect() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatalln("Error with connection to database: ", err.Error())
	}

	fmt.Println("Success - database is connected")
	defer db.Close()

	insert, errquery := db.Query("CREATE TABLE IF NOT EXISTS user ( ID INT PRIMARY KEY AUTO_INCREMENT, username VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL);")
	defer insert.Close()
	insert_blog, errquery_blog := db.Query("CREATE TABLE IF NOT EXISTS blogs ( ID INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(255) NOT NULL UNIQUE, mess VARCHAR(255) NOT NULL);")
	defer insert_blog.Close()

	if errquery != nil {
		log.Fatal(errquery.Error())
	}
	if errquery_blog != nil {
		log.Fatal(errquery_blog.Error())
	}

}

func main() {
	databaseConnect()
	routerStart()
}
