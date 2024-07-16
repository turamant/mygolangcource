// 1. Home (Hello World)
// Первый шаг - это создание самого простого веб-сервера,
//  который выводит "Hello, World!".

package main

import (
	"fmt"
	"net/http"
)

const port=":8001"

func routeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello world!")
}


func main(){
	http.HandleFunc("/", routeHandler)
	fmt.Printf("Start server on port:%s\n", port)
	http.ListenAndServe(port, nil)
}