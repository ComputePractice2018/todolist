package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/todolist/backend/server"
)

func main() {
	/*	var name = flag.String("name", "Mikhail", "имя участника команды")
		var team = flag.String("team", "Todolist", "название команды")*/
	//flag.Parse()

	router := mux.NewRouter()
	//роутинг списка задач
	router.HandleFunc("/api/todolist/task/", server.GetList).Methods("GET")
	router.HandleFunc("/api/todolist/task/", server.AddList).Methods("POST")
	router.HandleFunc("/api/todolist/task/{id}", server.EditList).Methods("PUT")
	router.HandleFunc("/api/todolist/task/{id}", server.DeleteList).Methods("DELETE")

	//роутинг подзадач в списке
	router.HandleFunc("/api/todolist/field/{idRef}", server.AddTask).Methods("POST")
	router.HandleFunc("/api/todolist/field/{idRef}/{idField}", server.EditTask).Methods("PUT")
	router.HandleFunc("/api/todolist/field/{idField}", server.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/todolist/field/success/{idRef}/{idField}", server.SuccessTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
