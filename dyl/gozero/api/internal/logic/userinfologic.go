package logic

import (
	"context"
	"fmt"
	"gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gozero/api/internal/svc"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	user2, err := l.svcCtx.UsersModel.FindOne(context.Background(), 2)
	fmt.Println(user, user2, err)

	return &types.UserInfoResponse{
		Id:       uint(user2.Id),
		UserName: user2.Username,
	}, nil
}
