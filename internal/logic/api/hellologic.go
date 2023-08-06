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
	res, result, err := l.svcCtx.Cbn.EnforceEx(req.Sub, req.Obj, req.Act)
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}

	policy := l.svcCtx.Cbn.GetPolicy()
	//policy := l.svcCtx.Cbn.GetFilteredPolicy(1, "")
	l.Logger.Infof("policy: %v", policy)
	resp.Message = fmt.Sprintf("%v", res)
	l.Logger.Infof("hello: %v result:%v err:%v", res, result, err)

	return resp, nil
}
