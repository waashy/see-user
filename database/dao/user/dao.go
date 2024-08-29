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
	Create() error
	Get() error
	List() error
	Update() error
	Delete() error
}

func init() {}

func NewUserDao(db *database.Database, logger *log.Logger) (UserDao, error) {

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

func (ud *userDao) Create() error { return nil }

func (ud *userDao) Get() error { return nil }

func (ud *userDao) List() error { return nil }

func (ud *userDao) Update() error { return nil }

func (ud *userDao) Delete() error { return nil }
