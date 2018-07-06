package utils

import "fmt"

//GetNameAndTeam возвращает приветствие участнику команды
func GetNameAndTeam(Name, Team string) string {
	return fmt.Sprintf("Hello, %s, from %s team", Name, Team)
}