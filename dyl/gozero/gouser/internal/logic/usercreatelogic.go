package logic

import (
	"context"
	"fmt"
	"gozero/gouser/internal/svc"
	"gozero/gouser/types/gouser"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *gouser.UserCreateRequest) (*gouser.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Username, in.Password)
	return &gouser.UserCreateResponse{}, nil
}
