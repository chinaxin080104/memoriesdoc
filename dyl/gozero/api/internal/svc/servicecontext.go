package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero/api/internal/config"
	"gozero/api/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//获取mysql链接
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserModel(mysqlConn),
	}
}
