package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/todolist/backend/data"
)

//GetList обрабатывает запросы на получение списка задач
func GetList(w http.ResponseWriter, r *http.Request) {
	BinaryData, err := json.Marshal(data.ListResponce)
	if err != nil {
		w.Header().Add("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON marshling error %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(BinaryData)
	if err != nil {
		log.Printf("Handler write error %v", err)
	}
}
