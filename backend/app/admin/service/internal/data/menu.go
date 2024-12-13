package data

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	timeutil "github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"
	"kratos-monolithic-demo/app/admin/service/internal/data/ent/menu"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	systemV1 "kratos-monolithic-demo/api/gen/go/system/service/v1"
)

type MenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) *MenuRepo {
	l := log.NewHelper(log.With(logger, "module", "menu/repo/admin-service"))
	return &MenuRepo{
		data: data,
		log:  l,
	}
}

func (r *MenuRepo) convertMenuTypeToProto(in *menu.Type) *systemV1.MenuType {
	if in == nil {
		return nil
	}
	find, ok := systemV1.MenuType_value[string(*in)]
	if !ok {
		return nil
	}
	return (*systemV1.MenuType)(trans.Ptr(find))
}
func (r *MenuRepo) convertMenuTypeToEnt(in *systemV1.MenuType) *menu.Type {
	if in == nil {
		return nil
	}
	find, ok := systemV1.MenuType_name[int32(*in)]
	if !ok {
		return nil
	}
	return (*menu.Type)(trans.Ptr(find))
}

func (r *MenuRepo) convertEntToProto(in *ent.Menu) *systemV1.Menu {
	if in == nil {
		return nil
	}

	return &systemV1.Menu{
		Id:         trans.Ptr(in.ID),
		ParentId:   in.ParentID,
		Type:       r.convertMenuTypeToProto(in.Type),
		Path:       in.Path,
		Redirect:   in.Redirect,
		Alias:      in.Alias,
		Name:       in.Name,
		Component:  in.Component,
		Meta:       in.Meta,
		Status:     (*string)(in.Status),
		CreateTime: timeutil.TimeToTimestamppb(in.CreateTime),
		UpdateTime: timeutil.TimeToTimestamppb(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimestamppb(in.DeleteTime),
	}
}

func (r *MenuRepo) travelChild(nodes []*systemV1.Menu, node *systemV1.Menu) bool {
	if nodes == nil {
		return false
	}

	if node.ParentId == nil {
		nodes = append(nodes, node)
		return true
	}

	for _, n := range nodes {
		if node.ParentId == nil {
			continue
		}

		if n.GetId() == node.GetParentId() {
			n.Children = append(n.Children, node)
			return true
		} else {
			if r.travelChild(n.Children, node) {
				return true
			}
		}
	}
	return false
}

func (r *MenuRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().Menu.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
	}

	return count, err
}

func (r *MenuRepo) List(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListMenuResponse, error) {
	builder := r.data.db.Client().Menu.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), menu.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("解析SELECT条件发生错误[%s]", err.Error())
		return nil, err
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, err
	}

	items := make([]*systemV1.Menu, 0, len(results))
	for _, m := range results {
		if m.ParentID == nil {
			item := r.convertEntToProto(m)
			items = append(items, item)
		}
	}
	for _, m := range results {
		if m.ParentID != nil {
			item := r.convertEntToProto(m)
			r.travelChild(items, item)
		}
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &systemV1.ListMenuResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *MenuRepo) IsExist(ctx context.Context, id int32) (bool, error) {
	return r.data.db.Client().Menu.Query().
		Where(menu.IDEQ(id)).
		Exist(ctx)
}

func (r *MenuRepo) Get(ctx context.Context, req *systemV1.GetMenuRequest) (*systemV1.Menu, error) {
	ret, err := r.data.db.Client().Menu.Get(ctx, req.GetId())
	if err != nil && !ent.IsNotFound(err) {
		r.log.Errorf("query one data failed: %s", err.Error())
		return nil, err
	}

	return r.convertEntToProto(ret), err
}

func (r *MenuRepo) Create(ctx context.Context, req *systemV1.CreateMenuRequest) error {
	if req.Menu == nil {
		return errors.New("invalid request")
	}

	builder := r.data.db.Client().Menu.Create().
		SetNillableParentID(req.Menu.ParentId).
		SetNillableType(r.convertMenuTypeToEnt(req.Menu.Type)).
		SetNillablePath(req.Menu.Path).
		SetNillableRedirect(req.Menu.Redirect).
		SetNillableAlias(req.Menu.Alias).
		SetNillableName(req.Menu.Name).
		SetNillableComponent(req.Menu.Component).
		SetNillableStatus((*menu.Status)(req.Menu.Status))

	if req.Menu.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	} else {
		builder.SetCreateTime(*timeutil.TimestamppbToTime(req.Menu.CreateTime))
	}

	if req.Menu.Meta != nil {
		builder.SetMeta(req.Menu.Meta)
	}

	if req.Menu.Id != nil {
		builder.SetID(req.Menu.GetId())
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *MenuRepo) Update(ctx context.Context, req *systemV1.UpdateMenuRequest) error {
	if req.Menu == nil {
		return errors.New("invalid request")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetMenu().GetId())
		if err != nil {
			return err
		}
		if !exist {
			return r.Create(ctx, &systemV1.CreateMenuRequest{Menu: req.Menu, OperatorId: req.OperatorId})
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Menu) {
			return errors.New("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetMenu(), req.UpdateMask.GetPaths())
	}

	r.log.Infof("UPDATE: [%v] [%v]", req.Menu, req.Menu.Meta)

	builder := r.data.db.Client().Menu.UpdateOneID(req.Menu.GetId()).
		SetNillableParentID(req.Menu.ParentId).
		SetNillableType(r.convertMenuTypeToEnt(req.Menu.Type)).
		SetNillablePath(req.Menu.Path).
		SetNillableRedirect(req.Menu.Redirect).
		SetNillableAlias(req.Menu.Alias).
		SetNillableName(req.Menu.Name).
		SetNillableComponent(req.Menu.Component).
		SetNillableStatus((*menu.Status)(req.Menu.Status))

	if req.Menu.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	} else {
		builder.SetUpdateTime(*timeutil.TimestamppbToTime(req.Menu.UpdateTime))
	}

	if req.Menu.Meta != nil {
		r.updateMetaField(builder, req.Menu.Meta)
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Menu, req.GetUpdateMask().GetPaths())
		_, nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	err := builder.Exec(ctx)
	if err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return err
	}

	return nil
}

func (r *MenuRepo) updateMetaField(builder *ent.MenuUpdateOne, meta *systemV1.RouteMeta) {
	builder.SetMeta(meta)

	builder.Modify(func(u *sql.UpdateBuilder) {
		//menu.FieldMeta.Update(u, sqljson.Value(meta))
	})
}

func (r *MenuRepo) Delete(ctx context.Context, req *systemV1.DeleteMenuRequest) (bool, error) {
	err := r.data.db.Client().Menu.
		DeleteOneID(req.GetId()).
		Exec(ctx)
	if err != nil {
		r.log.Errorf("delete one data failed: %s", err.Error())
	}

	return err == nil, err
}
