package user

import (
	log "log/slog"

	database "github.com/waashy/utils/database"
)

type userDao struct {
	db     *database.Database
	logger *log.Logger
}

type UserDao interface {
	Create()
	Get()
	List()
	Update()
	Delete()
}

func NewUserDao(db *database.Database, logger *log.Logger) (UserDao, error) {

	// auto migrate user

	if err := db.Instance.AutoMigrate(&User{}); err != nil {
		return nil, err
	}

	return &userDao{
		db:     db,
		logger: logger,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////// IMPLEMENTATIONS //////////////////////////////////////////////

func (ud *userDao) Create() {}

func (ud *userDao) Get() {}

func (ud *userDao) List() {}

func (ud *userDao) Update() {}

func (ud *userDao) Delete() {}
