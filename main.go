package main

import (
	"fmt"
	"net/http"
)

func main () {
	server := http.FileServer(http.Dir("files"))
	http.Handle("/", server)

	port := "8080"
	fmt.Println(("Serving on http://localhost:" + port))

	err := http.ListenAndServe(":"+port,nil)
	if err != nil {
		fmt.Println(("Error starting server: " + err.Error()))
	}

	
}