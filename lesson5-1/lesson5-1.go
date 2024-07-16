// 5-1. Работа с Базой данных и шаблонами
// Подключение к базе данных MySQL и выполнение простых операций INSERT and SELECT.
// renderTamplate 

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func connectDB() (*sql.DB, error) {
	dsn := "root:password@tcp(172.17.0.3:3306)/mygophercourse"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id INT AUTO_INCREMENT PRIMARY KEY, message VARCHAR(255))")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func homeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			message := r.FormValue("message")
			_, err := db.Exec("INSERT INTO messages (message) VALUES (?)", message)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		rows, err := db.Query("SELECT message FROM messages")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var messages []string
		for rows.Next() {
			var msg string
			err = rows.Scan(&msg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			messages = append(messages, msg)
		}

		data := struct {
			Messages []string
		}{
			Messages: messages,
		}
		renderTemplate(w, "index_bd.html", data)
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler(db)).Methods("GET", "POST")

	fmt.Println("Server is listening on port 8000...")
	http.ListenAndServe(":8000", r)
}
