// 3. Routing (using gorilla/mux)
// Теперь добавим поддержку маршрутизации с помощью gorilla/mux.

// Установка зависимости gorilla/mux
// Вам нужно установить пакет gorilla/mux, выполнив команду:

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello index page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is about page")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")

	server := &http.Server{
		Addr: ":8000",
		Handler: r,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	server.ListenAndServe()
}