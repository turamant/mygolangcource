package main

import (
	"fmt"
	"net/http"
	"time"
)


func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is root page")
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	
	server := &http.Server{
		Addr: ":8000",
		Handler: mux,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	fmt.Printf("Start server on port:%s\n", server.Addr)
	server.ListenAndServe()
	
}