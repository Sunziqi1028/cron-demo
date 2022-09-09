package logic

import (
	"context"
	"cron/service/task/api/internal/http"
	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskExcelToEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskExcelToEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskExcelToEmailLogic {
	return &TaskExcelToEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskExcelToEmailLogic) TaskExcelToEmail(req *types.ExcelMailRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	RunTask(req, l.svcCtx)
	return
}

func RunTask(request interface{}, svcCtx *svc.ServiceContext) {
	http.SendMail(svcCtx.Config.Email.Host, svcCtx.Config.Email.UserName, svcCtx.Config.Email.Password, "", "", "", "", "")
}
