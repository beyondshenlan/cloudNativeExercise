//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"week04/internal/biz"
	"week04/internal/data"
)

//var UserStoreSet = wire.NewSet(biz.NewUserBiz, data.NewUserRepo)

func InitUserBiz(str string) *biz.UserBiz {

	wire.Build(biz.NewUserBiz, data.NewUserRepo, data.NewDB)
	//wire.Build(UserSet)
	return &biz.UserBiz{}
}
