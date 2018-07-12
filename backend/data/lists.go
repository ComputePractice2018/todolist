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

//ListInterface - интерфейс для работы со списком задач
type ListInterface interface {
	GetList() []List
	AddList(list List) int
	EditList(list List, ID int) error
	DeleteList(ID int) error
	//CompleteList(ID int) error
}

//ListTask структура для списка задач
type ListTask struct {
	lists []List
}

//NewListTask конструктор списка задач
func NewListTask() *ListTask {
	return &ListTask{}
}

//GetList возвращает список задач
func (lt *ListTask) GetList() []List {
	return lt.lists
}

//AddList добавляет список задач и возвращает ID
func (lt *ListTask) AddList(list List) int {
	id := len(lt.lists) + 1
	list.ID = id
	lt.lists = append(lt.lists, list)
	return id
}

//EditList изменяет список задач с ID на list
func (lt *ListTask) EditList(list List, ID int) error {
	if ID < 1 || ID > len(lt.lists) {
		return fmt.Errorf("Incorrect ID")
	}
	lt.lists[ID-1] = list
	return nil
}

//DeleteList удаляет список задач по ID
func (lt *ListTask) DeleteList(ID int) error {
	if ID < 1 || ID > len(lt.lists) {
		return fmt.Errorf("Incorrect ID")
	}
	ID--
	lt.lists = append(lt.lists[:ID], lt.lists[ID+1:]...)

	return nil
}
