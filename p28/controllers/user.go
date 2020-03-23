package controllers

import (
	"context"
	"gomicro_note/p28/appInit"
	"gomicro_note/p28/models"
	"time"
)

type UserService struct {
}

func (*UserService) UserReg(ctx context.Context, req *models.UserModel, rsp *models.RegResponse) error {
	users := models.User{UserName: req.UserName, UserPwd: req.UserPwd, UserDate: time.Now()}
	err := appInit.GetDB().Create(&users).Error
	if err != nil {
		rsp.Message = err.Error()
		rsp.Status = "error"
	} else {
		rsp.Message = ""
		rsp.Status = "success"
	}

	return nil
}
