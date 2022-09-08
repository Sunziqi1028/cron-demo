package logic

import (
	"context"

	"cron/service/task/rpc/internal/svc"
	"cron/service/task/rpc/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *task.Req) (*task.TaskReply, error) {
	// todo: add your logic here and delete this line

	return &task.TaskReply{}, nil
}
