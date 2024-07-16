package main

import (
	"fmt"
	"net/http"
)

const port=":8000"

func routeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello this is root page")
}


func main(){
	http.HandleFunc("/", routeHandler)
	fmt.Printf("Start server on port:%s\n", port)
	http.ListenAndServe(port, nil)
}