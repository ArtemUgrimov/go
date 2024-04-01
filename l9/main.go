package main

import (
	"main/l9/pkg/school"
	"net/http"
)

// render html

func main() {
	school := school.EstablishSchool()
	http.HandleFunc("/", school.RootHandler)
	http.HandleFunc("/students", school.StudentPickerHandler)
	http.HandleFunc("/student/", school.GetStudentHandler)

	http.ListenAndServe(":8000", nil)
}
