package main

import (
	"main/l9/pkg/school"
	"net/http"
)

// render html

func main() {
	school := school.EstablishSchool()
	http.HandleFunc("/", school.RootHandler)
	http.HandleFunc("/students", school.Authorize(school.StudentPickerHandler, "/"))
	http.HandleFunc("GET /student/{id}", school.Authorize(school.GetStudentHandler, "/students"))

	http.ListenAndServe(":8000", nil)
}
