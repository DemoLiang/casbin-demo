package api

import (
	"context"
	"fmt"

	"github.com/DemoLiang/casbin-demo/internal/svc"
	"github.com/DemoLiang/casbin-demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{}
	res, err := l.svcCtx.Cbn.Enforce(req.Sub, req.Obj, req.Act)
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}
	resp.Message = fmt.Sprintf("%v", res)

	return resp, nil
}
