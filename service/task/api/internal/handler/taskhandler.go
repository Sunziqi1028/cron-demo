package handler

import (
	"net/http"

	"cron/service/task/api/internal/logic"
	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func taskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskLogic(r.Context(), svcCtx)
		resp, err := l.Task(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
