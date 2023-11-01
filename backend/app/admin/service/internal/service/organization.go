package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-monolithic-demo/app/admin/service/internal/data"

	"kratos-monolithic-demo/gen/api/go/common/pagination"
	v1 "kratos-monolithic-demo/gen/api/go/user/service/v1"
	v12 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type OrganizationService struct {
	v12.UnimplementedOrganizationServiceServer

	log *log.Helper

	uc *data.OrganizationRepo
}

func NewOrganizationService(uc *data.OrganizationRepo, logger log.Logger) *OrganizationService {
	l := log.NewHelper(log.With(logger, "module", "organization/service/admin-service"))
	return &OrganizationService{
		log: l,
		uc:  uc,
	}
}

func (s *OrganizationService) ListOrganization(ctx context.Context, req *pagination.PagingRequest) (*v1.ListOrganizationResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *OrganizationService) GetOrganization(ctx context.Context, req *v1.GetOrganizationRequest) (*v1.Organization, error) {
	return s.uc.Get(ctx, req)
}

func (s *OrganizationService) CreateOrganization(ctx context.Context, req *v1.CreateOrganizationRequest) (*v1.Organization, error) {
	return s.uc.Create(ctx, req)
}

func (s *OrganizationService) UpdateOrganization(ctx context.Context, req *v1.UpdateOrganizationRequest) (*v1.Organization, error) {
	return s.uc.Update(ctx, req)
}

func (s *OrganizationService) DeleteOrganization(ctx context.Context, req *v1.DeleteOrganizationRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
