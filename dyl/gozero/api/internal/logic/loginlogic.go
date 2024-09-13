package logic

import (
	"context"
	"gozero/api/internal/config"
	"gozero/api/utils/jwts"

	"gozero/api/internal/svc"
	"gozero/api/internal/types"

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
	AuthConfig := config.AuthConfig{}
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Username: "小疯子",
		Role:     1,
	}, AuthConfig.AccessSecret, AuthConfig.AccessExpire)
	if err != nil {
		return "", err
	}
	return token, err
}
