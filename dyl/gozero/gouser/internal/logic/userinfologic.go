package logic

import (
	"context"

	"gozero/gouser/internal/svc"
	"gozero/gouser/types/gouser"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *gouser.UserInfoRequest) (*gouser.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &gouser.UserInfoResponse{
		UserId:   in.UserId,
		Username: "英语六级",
	}, nil
}
