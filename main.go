package main
//import "fmt"
import (
	"net/http"
	//"io"
	"log"
	"fmt"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
var errdb error
db, errdb = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/users")
if errdb != nil {
    log.Fatal(errdb.Error())
}
defer db.Close()

insert, errquery := db.Query( "CREATE TABLE IF NOT EXISTS user ( ID INT PRIMARY KEY AUTO_INCREMENT, username VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(255) NOT NULL);")
insert2, errquery2 := db.Query( "CREATE TABLE IF NOT EXISTS blogs ( ID INT PRIMARY KEY AUTO_INCREMENT, title VARCHAR(255) NOT NULL UNIQUE, mess VARCHAR(255) NOT NULL);")

if errquery != nil  {
	log.Fatal(errquery.Error())
}
if errquery2 != nil {
	log.Fatal(errquery2.Error())
}

defer insert.Close()
defer insert2.Close()

fmt.Println("Success!")

http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w,r,"index.html")
})

http.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"login.html")
})
http.HandleFunc("/blog",func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"blog.html")
})
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

http.HandleFunc("/submit",submitHandle)
http.HandleFunc("/mess",submitMess)

log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func submitMess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Jeśli nie jest to żądanie POST, wyślij kod 405 Method Not Allowed
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form.", http.StatusInternalServerError)
		return
	}
	title   := r.FormValue("title")
	content := r.FormValue("content")

	_, errin := db.Exec("INSERT INTO blogs (title, mess) VALUES (?, ?)", title, content)
	if errin != nil {
		http.Error(w, "Error inserting new blog.", http.StatusInternalServerError)
		return
	}else
	{
		http.Redirect(w,r,"/blog",302)
	}

}


func submitHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Jeśli nie jest to żądanie POST, wyślij kod 405 Method Not Allowed
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form.", http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Użyj Exec zamiast Query dla operacji INSERT
	_, errin := db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	if errin != nil {
		http.Error(w, "Error inserting new user.", http.StatusInternalServerError)
		return
	}else
	{
		http.Redirect(w,r,"/",302)
	}


}