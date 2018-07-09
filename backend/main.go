package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/todolist/backend/data"
	"github.com/ComputePractice2018/todolist/backend/server"
)

func main() {
	/*	var name = flag.String("name", "Mikhail", "имя участника команды")
		var team = flag.String("team", "Todolist", "название команды")*/
	//flag.Parse()
	listTask := data.NewListTask()
	router := mux.NewRouter()
	//роутинг списка задач
	router.HandleFunc("/api/todolist/task/", server.GetList(listTask)).Methods("GET")
	router.HandleFunc("/api/todolist/task/", server.AddList(listTask)).Methods("POST")
	router.HandleFunc("/api/todolist/task/{id}", server.EditList(listTask)).Methods("PUT")
	router.HandleFunc("/api/todolist/task/{id}", server.DeleteList(listTask)).Methods("DELETE")
	router.HandleFunc("/api/todolist/task/complete/{id}", server.CompleteList(listTask)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
