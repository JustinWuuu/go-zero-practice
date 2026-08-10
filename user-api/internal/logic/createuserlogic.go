// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"user-rpc/user"

	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user.CreateUserReq{
		Name: req.Name,
	})	

	if err != nil {
		return nil, err
	}

	return &types.CreateUserResp{
		Id:   rpcResp.Id,
		Name: rpcResp.Name,
	}, nil
}
