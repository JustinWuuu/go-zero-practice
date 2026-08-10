package logic

import (
	"context"

	"order-rpc/internal/svc"
	"order-rpc/order"
	"user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoLogic {
	return &GetOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoLogic) GetOrderInfo(in *order.OrderInfoReq) (*order.OrderInfoResp, error) {
	userResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.UserInfoReq{
		Id: "123",
	})

	if err != nil {
		return nil, err
	}

	l.Logger.Infof("searched username: %s", userResp.Name)	

	return &order.OrderInfoResp{
		Id:      in.Id,
		Product: "test product",
		Amount:  100,
	}, nil
}
