package main

import "fmt"

//GetNameAndTeam возвращает приветствие участнику команды
func GetNameAndTeam(Name, Team string) string {
	return fmt.Sprintf("Hello, %s, from %s team", Name, Team)
}

func main() {
	fmt.Println("Hello World")
	fmt.Printf(GetNameAndTeam("Egor", "Todolist"))
}
