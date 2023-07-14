package backend

import (
	"encoding/json"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hibiken/asynq"

	"shop/api/backend"
	"shop/boot/job"
	"shop/internal/consts"
	"shop/internal/model"
	"shop/internal/service"

	"golang.org/x/net/context"
)

// Role 角色管理
var Role = cRole{}

type cRole struct{}

func (c *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.RoleId}, nil
}

func (c *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	out, err := service.Role().Update(ctx, model.RoleUpdateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &backend.RoleUpdateRes{
		RoleId: out.RoleId,
	}, nil
}

func (c *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &backend.RoleDeleteRes{}, nil
}

func (c *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	return &backend.RoleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// Assign 分配权限
func (c *cRole) Assign(ctx context.Context, req *backend.AddRolePermissionReq) (res *backend.AddRolePermissionRes, err error) {
	err = service.Role().AssignPermission(ctx, model.RoleAddPermissionInput{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AddRolePermissionRes{}, nil
}

func (c *cRole) CancelAssign(ctx context.Context, req *backend.DeletePermissionReq) (res *backend.DeletePermissionRes, err error) {
	err = service.Role().CancelAssignPermission(ctx, model.RoleDeletePermissionInput{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		return nil, err
	}
	return &backend.DeletePermissionRes{}, nil
}

func (c *cRole) GetAssignList(ctx context.Context, req *backend.RoleGetPermissionReq) (res *backend.RoleGetPermissionRes, err error) {
	// 1. 测试 Asynq 消息队列
	payload, err := json.Marshal(job.TestQueuePayload{
		Name: "test",
		Time: gtime.Now(),
	})

	if err != nil {
		return nil, err
	}
	_, err = job.AsynqClient.Enqueue(
		asynq.NewTask(consts.JobTestQueue, payload),
		asynq.MaxRetry(3), // 最大重试次数
	)
	if err != nil {
		return nil, err
	}

	// 1. 获取角色信息
	roleInfo, err := service.Role().GetPermissionList(ctx, model.RoleGetPermissionListInput{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	// 赋值
	return &backend.RoleGetPermissionRes{
		Role:               roleInfo.Role,
		RolePermissionList: roleInfo.RolePermissionList,
	}, nil

}
