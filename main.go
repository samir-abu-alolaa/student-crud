package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/Get", getStudentsHandler)
	http.HandleFunc("/Put", putStudentHandler)
	http.HandleFunc("/Delete", deleteStudentHandler)
	http.HandleFunc("/UpdateName", updateStudentNameHandler)
	http.HandleFunc("/Get-Student", getStudentByNameHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
