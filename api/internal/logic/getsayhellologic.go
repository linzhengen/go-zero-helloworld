package logic

import (
	"context"
	"helloworld/rpc/helloworld"

	"helloworld/api/internal/svc"
	"helloworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSayHelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSayHelloLogic {
	return &GetSayHelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSayHelloLogic) GetSayHello(req *types.Request) (resp *types.Response, err error) {
	result, err := l.svcCtx.HelloWorldGreeterClient.SayHello(l.ctx, &helloworld.HelloRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("say hello: %s", req.Name)
	return &types.Response{Message: result.Message}, nil
}
