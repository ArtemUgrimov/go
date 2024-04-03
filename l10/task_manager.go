package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type TaskManager struct {
	tasks map[int]Task
}

func (tm *TaskManager) getTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if len(tm.tasks) > 0 {
		json.NewEncoder(w).Encode(tm.tasks)
	} else {
		w.Write([]byte("{}"))
	}
}

func (tm *TaskManager) getTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HttpError(w, "Cannot parse ID", http.StatusBadRequest)
		return
	}

	task, ok := tm.tasks[id]
	if !ok {
		HttpError(w, "Cannot find task with given ID", http.StatusNotFound)
		return
	}
	data, error := json.Marshal(task)
	if error != nil {
		HttpError(w, error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (tm *TaskManager) delTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HttpError(w, "Cannot parse ID", http.StatusBadRequest)
		return
	}

	_, ok := tm.tasks[id]
	if !ok {
		HttpError(w, "Cannot find task with given ID", http.StatusNotFound)
		return
	}
	delete(tm.tasks, id)
	w.WriteHeader(http.StatusOK)
}

func (tm *TaskManager) putTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		HttpError(w, "Cannot parse ID", http.StatusBadRequest)
		return
	}

	task, ok := tm.tasks[id]
	if !ok {
		HttpError(w, "Cannot find task with given ID", http.StatusNotFound)
		return
	}

	// do modifications with the task
	// for example parse new task and change original values
	var newTask Task
	err = json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (Task{}) == newTask {
		HttpError(w, "Cannot parse task", http.StatusBadRequest)
		return
	}
	// ignore ID sent by the user since we won't let changing it
	newTask.Id = task.Id
	tm.tasks[id] = newTask

	data, error := json.Marshal(newTask)
	if error != nil {
		HttpError(w, error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (tm *TaskManager) postTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (Task{}) == newTask {
		HttpError(w, "Cannot parse task", http.StatusBadRequest)
		return
	}
	_, ok := tm.tasks[newTask.Id]
	if ok {
		HttpError(w, "There is already a task with such ID", http.StatusBadRequest)
		return
	}

	tm.tasks[newTask.Id] = newTask
	json.NewEncoder(w).Encode(newTask)
}

func (tm *TaskManager) getRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /tasks", tm.getTasks)
	router.HandleFunc("GET /tasks/{id}", tm.getTaskByID)
	router.HandleFunc("DELETE /tasks/{id}", tm.delTaskByID)
	router.HandleFunc("PUT /tasks/{id}", tm.putTaskByID)
	router.HandleFunc("POST /tasks", tm.postTask)

	return router
}

func HttpError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"status": "error","message": "%s"}`, message)))
}

func NewTaskManager() *TaskManager {
	tm := &TaskManager{}
	tm.tasks = make(map[int]Task)

	return tm
}
