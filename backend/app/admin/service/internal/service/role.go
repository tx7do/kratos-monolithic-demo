package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-monolithic-demo/app/admin/service/internal/data"

	"kratos-monolithic-demo/gen/api/go/common/pagination"
	v12 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type RoleService struct {
	v12.UnimplementedRoleServiceServer

	log *log.Helper

	uc *data.RoleRepo
}

func NewRoleService(uc *data.RoleRepo, logger log.Logger) *RoleService {
	l := log.NewHelper(log.With(logger, "module", "role/service/admin-service"))
	return &RoleService{
		log: l,
		uc:  uc,
	}
}

func (s *RoleService) ListRole(ctx context.Context, req *pagination.PagingRequest) (*v12.ListRoleResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *RoleService) GetRole(ctx context.Context, req *v12.GetRoleRequest) (*v12.Role, error) {
	return s.uc.Get(ctx, req)
}

func (s *RoleService) CreateRole(ctx context.Context, req *v12.CreateRoleRequest) (*v12.Role, error) {
	return s.uc.Create(ctx, req)
}

func (s *RoleService) UpdateRole(ctx context.Context, req *v12.UpdateRoleRequest) (*v12.Role, error) {
	return s.uc.Update(ctx, req)
}

func (s *RoleService) DeleteRole(ctx context.Context, req *v12.DeleteRoleRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
