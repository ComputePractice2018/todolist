package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/todolist/backend/server"
)

func main() {
	/*	var name = flag.String("name", "Mikhail", "имя участника команды")
		var team = flag.String("team", "Todolist", "название команды")*/
	//flag.Parse()

	http.HandleFunc("/api/task/getList", server.GetList)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
