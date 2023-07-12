package role

import (
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"

	"shop/internal/dao"
	"shop/internal/model"
	"shop/internal/model/entity"
	"shop/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateUpdateOutput, err error) {
	// TODO: 判断角色是否存在

	// 插入数据返回id
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateUpdateOutput{RoleId: uint(lastInsertID)}, err
}

func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) (out model.RoleCreateUpdateOutput, err error) {
	// TODO: 判断角色是否存在

	// 更新数据
	_, err = dao.RoleInfo.Ctx(ctx).
		Data(in).
		OmitEmpty().
		Where(dao.RoleInfo.Columns().Id, in.Id).
		Update()
	if err != nil {
		return out, err
	}
	return model.RoleCreateUpdateOutput{RoleId: in.Id}, err
}

func (s *sRole) Delete(ctx context.Context, id uint) (err error) {
	// 删除数据
	_, err = dao.RoleInfo.Ctx(ctx).
		Where(dao.RoleInfo.Columns().Id, id).
		Unscoped(). // Unscoped() 会真实删除数据，而不是软删除
		Delete()

	return err
}

// GetList 查询内容列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.RoleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// 不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// AssignPermission 给角色分配权限
func (s *sRole) AssignPermission(ctx context.Context, in model.RoleAddPermissionInput) (err error) {
	List := g.List{}
	for _, v := range in.PermissionIds {
		List = append(List, g.Map{
			dao.RolePermissionInfo.Columns().RoleId:       in.RoleId,
			dao.RolePermissionInfo.Columns().PermissionId: v,
		})
	}
	_, err = dao.RolePermissionInfo.Ctx(ctx).Data(List).Insert()
	return err
}

// CancelAssignPermission 取消角色权限
func (s *sRole) CancelAssignPermission(ctx context.Context, in model.RoleDeletePermissionInput) (err error) {
	// return gerror.NewCode(response.CodeNotFound)
	_, _ = dao.RolePermissionInfo.Ctx(ctx).
		Where(dao.RolePermissionInfo.Columns().RoleId, in.RoleId).
		WhereIn(dao.RolePermissionInfo.Columns().PermissionId, in.PermissionIds).
		Delete()
	return nil
}
