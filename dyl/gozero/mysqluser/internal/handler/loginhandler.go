package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero/api/common/response"
	"gozero/mysqluser/internal/logic"
	"gozero/mysqluser/internal/svc"
	"gozero/mysqluser/internal/types"
	"net/http"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)

	}
}
