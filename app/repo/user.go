package repo

import (
	"echo-basic/app/model"
	"echo-basic/config"
)

var (
	UserRepo = new(userRepo)
)

type userRepo struct{}

func (userRepo) FindByName(name string) (*model.User, error) {
	db := config.NewDb()

	user := new(model.User)
	return user, db.Where("username = ?", name).First(user).Error
}
