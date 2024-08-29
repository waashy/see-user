package user

import (
	log "log/slog"

	userdao "github.com/waashy/see-user/database/dao/user"
)

type userService struct {
	dao    userdao.UserDao
	logger *log.Logger
}

type UserService interface {
	Start() error
	Stop() error

	Create() error
	Get() error
	List() error
	Update() error
	Delete() error
}

func NewUserService(us userdao.UserDao, logger *log.Logger) (UserService, error) {
	return &userService{
		dao:    us,
		logger: logger,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////// PROCESSESS /////////////////////////////////////////////////

func (us *userService) Start() error { return nil }

func (us *userService) Stop() error { return nil }

////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////// IMPLEMENTATIONS //////////////////////////////////////////////

func (us *userService) Create() error { return nil }

func (us *userService) Get() error { return nil }

func (us *userService) List() error { return nil }

func (us *userService) Update() error { return nil }

func (us *userService) Delete() error { return nil }
