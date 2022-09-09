package logic

import (
	"context"

	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskMessageToEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskMessageToEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskMessageToEmailLogic {
	return &TaskMessageToEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskMessageToEmailLogic) TaskMessageToEmail(req *types.MessageMailRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
