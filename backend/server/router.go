package server

import (
	"github.com/ComputePractice2018/todolist/backend/data"
	"github.com/gorilla/mux"
)

//NewRouter создает новый роутер
func NewRouter(listTask data.ListInterface) *mux.Router {
	router := mux.NewRouter()
	//роутинг списка задач
	router.HandleFunc("/api/todolist/task/", GetList(listTask)).Methods("GET")
	router.HandleFunc("/api/todolist/task/", AddList(listTask)).Methods("POST")
	router.HandleFunc("/api/todolist/task/{id}", EditList(listTask)).Methods("PUT")
	router.HandleFunc("/api/todolist/task/{id}", DeleteList(listTask)).Methods("DELETE")
	//router.HandleFunc("/api/todolist/task/complete/{id}", CompleteList(listTask)).Methods("PUT")
	return router
}
