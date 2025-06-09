package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/students", studentsHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
