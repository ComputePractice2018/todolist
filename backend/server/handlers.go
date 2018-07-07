package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/todolist/backend/data"
)

//ListHandler обрабатывает все запросы к api/task/getList
func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetList(w, r)
		return
	}
	if r.Method == "POST" {
		AddList(w, r)
		return
	}
	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Printf(message)
}

//GetList обрабатывает запросы на получение списка задач
func GetList(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.ListResponce)
	if err != nil {
		message := fmt.Sprintf("JSON marshling error %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Printf(message)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Printf(message)
	}
}

//AddList обрабатывает POST запрос
func AddList(w http.ResponseWriter, r *http.Request) {
	var list data.List
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", list)
	w.WriteHeader(http.StatusCreated)
}
