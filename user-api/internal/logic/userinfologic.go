// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"

	"user-api/internal/svc"
	"user-api/internal/types"
	"user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		Id:   rpcResp.Id,
		Name: rpcResp.Name,
	}, nil
}
