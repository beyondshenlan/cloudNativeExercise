package main

import (
	"database/sql"
	"errors"
	"fmt"
	pkg_errors "github.com/pkg/errors"
)

var NotFound = errors.New("not found data")
func Dao(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		return pkg_errors.Wrapf(NotFound,fmt.Sprintf("data not found ,sql:%s",query))
	}
	if err != nil{
		return pkg_errors.Wrapf(err,fmt.Sprintf("db system error ,sql:%s",query))
	}
	return nil
}

func Biz() error {
	err :=  Dao("")
	if errors.Is(err,NotFound){
		//业务处理
		return nil
	}
	if err != nil{
		return err
	}
	return nil
}
