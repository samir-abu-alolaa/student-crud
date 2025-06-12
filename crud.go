package main

import "database/sql"

func createStudent(Name, Email string, Age int) error {
	query := "INSERT INTO students (name, age, email) VALUES (?, ?, ?)"
	_, err := db.Exec(query, Name, Age, Email)
	return err
}

func getStudents() ([]Student, error) {
	rows, err := db.Query("SELECT id, name, age, email FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Email); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func getStudentByName(name string) (*Student, error) {
	query := "SELECT id, name, age, email FROM students WHERE name = ? LIMIT 1"
	row := db.QueryRow(query, name)

	var s Student
	err := row.Scan(&s.ID, &s.Name, &s.Age, &s.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func updateStudent(name, newName string) error {
	query := "UPDATE students SET name=? WHERE name=?"
	_, err := db.Exec(query, newName, name)
	return err
}

func deleteStudent(id int) error {
	query := "DELETE FROM students WHERE id=?"
	_, err := db.Exec(query, id)
	return err
}
