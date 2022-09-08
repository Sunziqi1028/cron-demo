package logic

import (
	"context"

	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogic {
	return &TaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskLogic) Task(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
