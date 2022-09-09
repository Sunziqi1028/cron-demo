package handler

import (
	"net/http"

	"cron/service/task/api/internal/logic"
	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func taskMessageToWechatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageWechatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTaskMessageToWechatLogic(r.Context(), svcCtx)
		resp, err := l.TaskMessageToWechat(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
