// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"

	"order-rpc/order"
	"user-api/internal/svc"
	"user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	rpcResp, err := l.svcCtx.OrderRpc.GetOrderInfo(l.ctx, &order.OrderInfoReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.OrderInfoResp{
		Id:    rpcResp.Id,
		Product: rpcResp.Product,
		Amount: rpcResp.Amount,
	}, nil
}
