package logic

import (
	"context"
	"fmt"

	"cron/service/task/rpc/internal/svc"
	"cron/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskToMonitorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskToMonitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskToMonitorLogic {
	return &TaskToMonitorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskToMonitorLogic) TaskToMonitor(in *task.Req) (*task.TaskReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("RPC服务被调用:", in.Id)
	fmt.Println("重新执行任务:", in.Id)
	fmt.Println("执行任务完成，变更数据库状态", in.Id)
	return &task.TaskReply{Id: 1, Message: "任务再次执行成功！"}, nil
}
