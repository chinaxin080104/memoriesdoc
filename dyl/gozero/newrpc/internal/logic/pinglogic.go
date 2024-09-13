package logic

import (
	"context"

	"gozero/newrpc/internal/svc"
	"gozero/newrpc/newrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *newrpc.Request) (*newrpc.Response, error) {
	// todo: add your logic here and delete this line

	return &newrpc.Response{}, nil
}
