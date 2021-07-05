package test

import (
	"context"
	"gomicro_note/p23/models"
	"strconv"
)

// TestService 结构体
type TestService struct {
}

// Call 方法
func (*TestService) Call(ctx context.Context, req *models.TestRequest, rsp *models.TestResponse) error {
	rsp.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}
