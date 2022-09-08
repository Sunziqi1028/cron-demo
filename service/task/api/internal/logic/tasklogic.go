package logic

import (
	"context"
	"fmt"

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
	fmt.Printf("当前任务执行：%s, %s, %s", req.StartTime, req.Duration, req.Task)
	fmt.Println("当前任务执行失败，变更数据库任务执行状态")
	return
}
