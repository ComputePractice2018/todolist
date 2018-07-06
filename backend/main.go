package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/todolist/backend/utils"
)

func main() {
	var name = flag.String("name", "Mikhail", "имя участника команды")
	var team = flag.String("team", "Todolist", "название команды")
	flag.Parse()
	fmt.Printf(utils.GetNameAndTeam(*name, *team))
}
