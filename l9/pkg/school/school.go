package school

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type School struct {
	students map[int]Student
	teachers []Teacher
}

func EstablishSchool() School {
	s := School{}
	s.students = make(map[int]Student)
	s.teachers = make([]Teacher, 0)
	s.teachers = append(s.teachers, NewTeacher("Artem", "1", true))
	s.teachers = append(s.teachers, NewTeacher("Mykola", "1", false))

	s.students[0] = NewStudent(0)
	s.students[1] = NewStudent(1)
	s.students[2] = NewStudent(2)
	s.students[3] = NewStudent(3)

	return s
}

func (school *School) GetStudent(id int) (Student, bool) {
	s, ok := school.students[id]
	if ok {
		return s, true
	}
	return Student{}, false
}

func (school *School) authenticateTeacher(login string, password string) (*Teacher, bool) {
	for idx, t := range school.teachers {
		if t.Name == login && t.Password == password {
			return &school.teachers[idx], true
		}
	}
	return nil, false
}

// ================ HANDLERS ===============//
func (school *School) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/index.html")
		return
	} else if r.Method == http.MethodPost {
		login := r.FormValue("login")
		pass := r.FormValue("password")

		t, ok := school.authenticateTeacher(login, pass)
		if ok {
			token := strconv.Itoa(rand.Int())
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   token,
				Expires: time.Now().Add(time.Minute),
			})
			t.token = token
			http.Redirect(w, r, "/students", http.StatusSeeOther)
			return
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (school *School) StudentPickerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/studentPicker.html")
		return
	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			http.Redirect(w, r, fmt.Sprintf("/student/%d", id), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/students", http.StatusSeeOther)
		}
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func (school *School) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	autorized := false
	for _, t := range school.teachers {
		if t.token == cookie.Value {
			autorized = true
			break
		}
	}
	if !autorized {
		http.Redirect(w, r, "/students", http.StatusSeeOther)
		return
	}

	studentReg := regexp.MustCompile(`/student/[0-9]+$`)
	if !studentReg.MatchString(r.URL.Path) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	studentId, found := strings.CutPrefix(r.URL.Path, "/student/")
	id, err := strconv.Atoi(studentId)
	if found && err == nil {
		student, ok := school.GetStudent(id)
		if ok {
			res, err := json.Marshal(student)
			if err == nil {
				w.Write(res)
				return
			}
		}
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}
