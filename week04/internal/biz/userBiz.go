package biz

import "week04/internal/data"

type UserBiz struct {
	userRepo *data.UserRepo
}

func NewUserBiz(repo *data.UserRepo) *UserBiz {
	return &UserBiz{
		userRepo: repo,
	}
}
