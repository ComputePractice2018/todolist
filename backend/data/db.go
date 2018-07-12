package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//DBListTask структура БД
type DBListTask struct {
	db *gorm.DB
}

//NewDBListTask функция создания нового списка задач в БД
func NewDBListTask(connection string) (*DBListTask, error) {
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&List{})
	return &DBListTask{db: db}, db.Error
}

//Close функция для закрытия БД
func (cl *DBListTask) Close() {
	cl.db.Close()
}

//GetList функция для получения списка задач из БД
func (cl *DBListTask) GetList() []List {
	var lists []List
	cl.db.Find(&lists)
	return lists
}

//AddList функция для добавления нового списка задач в БД
func (cl *DBListTask) AddList(list List) int {
	cl.db.Create(&list)
	return list.ID
}

//EditList функция редактирования списка задач в БД
func (cl *DBListTask) EditList(list List, id int) error {
	var lists []List
	cl.db.Where("id = ?", id).Find(&lists)
	if len(lists) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}

	list.ID = lists[0].ID
	cl.db.Save(&list)
	return cl.db.Error
}

//DeleteList функция удаления списка задач из БД
func (cl *DBListTask) DeleteList(id int) error {
	var lists []List
	cl.db.Where("id = ?", id).Find(&lists)
	if len(lists) == 0 {
		return fmt.Errorf("can't find record #%d", id)
	}
	cl.db.Delete(&lists[0])
	return cl.db.Error
}
