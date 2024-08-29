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

func init() {}

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

func (us *userService) Create() error { return us.dao.Create() }

func (us *userService) Get() error { return us.dao.Get() }

func (us *userService) List() error { return us.dao.List() }

func (us *userService) Update() error { return us.dao.Update() }

func (us *userService) Delete() error { return us.dao.Delete() }
