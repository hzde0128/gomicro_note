package test

import (
	"context"
	"gomicro_note/p23/models"
	"strconv"
)

type TestService struct {

}

func(*TestService)Call(ctx context.Context,req *models.TestRequest,rsp *models.TestResponse) error{
	rsp.Data = "test"+ strconv.Itoa(int(req.Id))
	return nil
}