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

	log.Fatal(http.ListenAndServe(":8080", router))
}
