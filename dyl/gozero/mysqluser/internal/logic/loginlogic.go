package logic

import (
	"context"
	"errors"
	"gozero/mysqluser/models"

	"gozero/mysqluser/internal/svc"
	"gozero/mysqluser/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	// todo: add your logic here and delete this line
	var user models.User
	err = l.svcCtx.DB.Take(&user, "username = ? and password = ?", req.Username, req.Password).Error
	if err != nil {
		return "", errors.New("登录失败")
	}
	return user.Username, nil
}
