package data

import (
	"fmt"
)

//List Структура для хранения списка задач
type List struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

//ListResponce хранимый список задач
var lists []List

//GetList возвращает список задач
func GetList() []List {
	return lists
}

//AddList добавляет список задач и возвращает ID
func AddList(list List) int {
	ID := len(lists)
	lists = append(lists, list)
	return ID
}

//EditList изменяет список задач с ID на list
func EditList(list List, ID int) error {
	if ID < 0 || ID >= len(lists) {
		return fmt.Errorf("Incorrect ID")
	}
	lists[ID] = list
	return nil
}

//DeleteList удаляет список задач по ID
func DeleteList(ID int) error {
	if ID < 0 || ID >= len(lists) {
		return fmt.Errorf("Incorrect ID")
	}
	lists = append(lists[:ID], lists[ID+1:]...)
	return nil
}

//CompleteList изменяет список задач с ID на list
func CompleteList(ID int) error {
	if ID < 0 || ID >= len(lists) {
		return fmt.Errorf("Incorrect ID")
	}
	switch lists[ID].Success {
	case true:
		lists[ID].Success = false
	case false:
		lists[ID].Success = true
	default:
		panic("Error")
	}
	return nil
}
