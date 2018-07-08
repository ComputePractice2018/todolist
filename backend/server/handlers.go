package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/todolist/backend/data"
	"github.com/gorilla/mux"
)

//GetList обрабатывает запросы на получение списка задач
func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetList())
	if err != nil {
		message := fmt.Sprintf("JSON to encode data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Printf(message)
		return
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
	id := data.AddList(list)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

//EditList обрабатывает PUT запрос
func EditList(w http.ResponseWriter, r *http.Request) {
	var list data.List
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		message := fmt.Sprintf("Unable to decode PUT data %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	err = data.EditList(list, id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)
}

//DeleteList обрабатывает DELETE запрос
func DeleteList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	err = data.DeleteList(id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}

func AddTask(w http.ResponseWriter, r *http.Request) {

}

func EditTask(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

func SuccessTask(w http.ResponseWriter, r *http.Request) {

}
