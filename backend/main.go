package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/todolist/backend/data"
	"github.com/ComputePractice2018/todolist/backend/server"
)

func main() {
	/*	var name = flag.String("name", "Mikhail", "имя участника команды")
		var team = flag.String("team", "Todolist", "название команды")*/
	//flag.Parse()
	listTask := data.NewListTask()
	router := server.NewRouter(listTask)
	log.Fatal(http.ListenAndServe(":8080", router))
}
