package controllers

import (
	"context"
	"gomicro_note/p27/models"
)

type UserService struct {
}

func (*UserService) UserReg(ctx context.Context, req *models.UserModel, rsp *models.RegResponse) error {
	rsp.Message = ""
	rsp.Status = "success"
	return nil
}
