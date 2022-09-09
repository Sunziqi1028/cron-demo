package logic

import (
	"context"

	"cron/service/task/api/internal/svc"
	"cron/service/task/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskMessageToWechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskMessageToWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskMessageToWechatLogic {
	return &TaskMessageToWechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskMessageToWechatLogic) TaskMessageToWechat(req *types.MessageWechatRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
