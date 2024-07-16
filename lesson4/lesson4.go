// 4. MySQL Database
// Подключение к базе данных MySQL и выполнение простых операций.

// Установка зависимости go-sql-driver/mysql

// docker run --name my-mysql -e MYSQL_ROOT_PASSWORD=password -d mysql:latest

// docker exec -it my-mysql mysql -uroot -ppassword -h "172.17.0.3"
// Как получить порт "172.17.0.3" ? 
// docker inspect my-mysql | grep IPAddress)


package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

func connectionDB() (*sql.DB, error){
	dsn := "root:password@tcp(172.17.0.3:3306)/mygophercourse"
	db, err := sql.Open("mysql", dsn)
	if err != nil{
		return nil, err
	}
	err = db.Ping()
	if err != nil{
		return nil, err
	}

	return db, nil
}

func homeHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		rows, err := db.Query("SELECT 'hello database!'")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next(){
			var msg string
			err = rows.Scan(&msg)
			if err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprint(w, msg)
		}
	}
}

func main()  {
	db, err := connectionDB()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler(db)).Methods("GET")
	
	server := http.Server{
		Addr: ":8000",
		Handler: r,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 15 * time.Second,
	}
	fmt.Printf("Start server on port:%s\n", server.Addr)
	server.ListenAndServe()
	
}