package logic

import (
	"GolangLearning/gozero/microhelloworld/mall/user/types/user"
	"context"

	"GolangLearning/gozero/microhelloworld/mall/order/internal/svc"
	"GolangLearning/gozero/microhelloworld/mall/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	userId := l.getOrderById(req.Id)
	userResponses, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderReply{
		Id:       req.Id,
		Name:     "zfy",
		Username: userResponses.Name,
	}, nil
}

func (l *GetOrderLogic) getOrderById(id string) string {
	return "1"
}
