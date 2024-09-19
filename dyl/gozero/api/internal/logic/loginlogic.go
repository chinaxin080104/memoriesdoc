package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gozero/api/internal/svc"
	"gozero/api/internal/types"
	"gozero/api/model"
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
	//var user model.UserModel
	//err = l.svcCtx.DB.Take(&user, "username = ? and password = ?", req.UserName, req.Password).Error
	//if err != nil {
	//	return "", errors.New("登录失败")
	//}
	//return "", errors.New("success")
	//操作mysql
	//添加
	l.svcCtx.UsersModel.Insert(context.Background(), &model.User{
		Username: "乔峰",
		Password: "123456",
	})
	//
	l.svcCtx.UsersModel.Insert(context.Background(), &model.User{
		Username: "慕容",
		Password: "123456",
	})
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(insert)
	//fmt.Println("登录成功")

	//查询
	//user, err := l.svcCtx.UsersModel.FindOne(context.Background(), 1)
	//user1, err := l.svcCtx.UsersModel.FindOneByUsername(context.Background(), "枫枫")
	//user2, err := l.svcCtx.UsersModel.FindOne(context.Background(), 2)
	//fmt.Println(user, user1, user2, err)

	//修改
	//l.svcCtx.UsersModel.Update(context.Background(), &model.User{Username: "倚天剑", Password: "987456", Id: 1})
	//l.svcCtx.UsersModel.Update(context.Background(), &model.User{Username: "屠龙刀", Password: "888888", Id: 2})
	//res, err := l.svcCtx.UsersModel.FindOne(context.Background(), 2)
	//fmt.Println(res, err)

	//删除   如果不存在 <nil> sql: no rows in result set
	//l.svcCtx.UsersModel.Delete(context.Background(), 6)
	//res3, err := l.svcCtx.UsersModel.FindOne(context.Background(), 6)
	//fmt.Println(res3, err)

	//return
	//获取config-的token 配置信息。
	//AuthConfig := config.AuthConfig{}

	////v1, 获取用户token -UserID,Username,Role 是写死的三个值。
	//token, err := jwts.GenToken(jwts.JwtPayLoad{
	//	UserID:   1,
	//	Username: "小疯子",
	//	Role:     1,
	//}, AuthConfig.AccessSecret, AuthConfig.AccessExpire)
	//
	//if err != nil {
	//	return "", err
	//}
	//
	return
}
