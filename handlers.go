package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func putStudentHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	email := r.URL.Query().Get("email")

	if strings.TrimSpace(name) == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(age) == "" {
		http.Error(w, "Missing age", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(email) == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	ag, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	err = createStudent(name, email, ag)
	if err != nil {
		http.Error(w, "Failed to create student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Student created successfully"))
}

func deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if strings.TrimSpace(id) == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}

	Id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = deleteStudent(Id)
	if err != nil {
		http.Error(w, "Failed to delete student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send 200 OK with confirmation message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
}

func updateStudentNameHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	name := r.URL.Query().Get("name")
	newName := r.URL.Query().Get("newName")

	// Validate input
	if strings.TrimSpace(name) == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(newName) == "" {
		http.Error(w, "Missing new name", http.StatusBadRequest)
		return
	}

	// Call the update function
	err := updateStudent(name, newName)
	if err != nil {
		http.Error(w, "Failed to update student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student name updated successfully"))
}

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := getStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(students)
}

func getStudentByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if strings.TrimSpace(name) == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return
	}
	student, err := getStudentByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if student == nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}
