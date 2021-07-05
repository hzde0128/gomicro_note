package controllers

import (
	"context"
	"gomicro_note/p28/appInit"
	"gomicro_note/p28/models"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
}

func (*UserService) UserReg(ctx context.Context, req *models.UserModel, rsp *models.RegResponse) error {
	users := models.UserModel{UserName: req.UserName, UserPwd: req.UserPwd, UserDate: timestamppb.Now()}
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
