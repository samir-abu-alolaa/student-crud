package main

import (
	"net/http"
)

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getStudentsHandler(w, r)
	case "POST":
		createStudentHandler(w, r)
	case "PUT":
		updateStudentHandler(w, r)
	case "DELETE":
		deleteStudentHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Implement getStudentsHandler, createStudentHandler, updateStudentHandler, deleteStudentHandler here...
