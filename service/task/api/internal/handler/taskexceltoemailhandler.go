package handler

import (
	"fmt"
	"net/http"

	"cron/service/task/api/internal/logic"
	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func taskExcelToEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExcelMailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println(req)
		l := logic.NewTaskExcelToEmailLogic(r.Context(), svcCtx)
		resp, err := l.TaskExcelToEmail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
