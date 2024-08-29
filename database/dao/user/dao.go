package user

import (
	database "github.com/waashy/utils/database"
)

type userDao struct {
	db *database.Database
}

type UserDao interface {
	Create()
	Get()
	List()
	Update()
	Delete()
}

func NewUserDao(db *database.Database) UserDao {
	return &userDao{
		db: db,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////// IMPLEMENTATIONS //////////////////////////////////////////////

func (ud userDao) Create() {}

func (ud userDao) Get() {}

func (ud userDao) List() {}

func (ud userDao) Update() {}

func (ud userDao) Delete() {}
