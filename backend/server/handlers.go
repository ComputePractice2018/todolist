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
func GetList(lt data.ListInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(lt.GetList())
		if err != nil {
			message := fmt.Sprintf("JSON to encode data: %v", err)
			http.Error(w, message, http.StatusInternalServerError)
			log.Printf(message)
			return
		}
	}
}

//AddList обрабатывает POST запрос
func AddList(lt data.ListInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var list data.List
		err := json.NewDecoder(r.Body).Decode(&list)
		if err != nil {
			message := fmt.Sprintf("Unable to decode POST data %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		id := lt.AddList(list)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditList обрабатывает PUT запрос
func EditList(lt data.ListInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
		err = lt.EditList(list, id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}

//DeleteList обрабатывает DELETE запрос
func DeleteList(lt data.ListInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		err = lt.DeleteList(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}

/*
//CompleteList обрабатывает PUT запрос на изменение выполненности задачи
func CompleteList(lt data.ListInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		err = lt.CompleteList(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)
	}
}
*/
