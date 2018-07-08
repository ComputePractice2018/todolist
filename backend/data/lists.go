package data

//List Структура для хранения списка задач
type List struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

//ListResponce хранимый список задач
var ListResponce = []List{List{
	ID:      1,
	Name:    "Название",
	Success: false,
}}
