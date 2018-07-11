package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/todolist/backend/data"
	"github.com/ComputePractice2018/todolist/backend/server"
)

func main() {
	connection := flag.String("connection", "todolist:SuperSecretPassword@tcp(db:3306)/todolist", "mysql connection string")
	flag.Parse()

	listTask, err := data.NewDBListTask(*connection)
	defer listTask.Close()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(listTask)
	log.Fatal(http.ListenAndServe(":8080", router))
}
