package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-monolithic-demo/app/admin/service/internal/data"

	"kratos-monolithic-demo/gen/api/go/common/pagination"
	v12 "kratos-monolithic-demo/gen/api/go/user/service/v1"
)

type PositionService struct {
	v12.UnimplementedPositionServiceServer

	log *log.Helper

	uc *data.PositionRepo
}

func NewPositionService(uc *data.PositionRepo, logger log.Logger) *PositionService {
	l := log.NewHelper(log.With(logger, "module", "position/service/admin-service"))
	return &PositionService{
		log: l,
		uc:  uc,
	}
}

func (s *PositionService) ListPosition(ctx context.Context, req *pagination.PagingRequest) (*v12.ListPositionResponse, error) {
	return s.uc.List(ctx, req)
}

func (s *PositionService) GetPosition(ctx context.Context, req *v12.GetPositionRequest) (*v12.Position, error) {
	return s.uc.Get(ctx, req)
}

func (s *PositionService) CreatePosition(ctx context.Context, req *v12.CreatePositionRequest) (*v12.Position, error) {
	return s.uc.Create(ctx, req)
}

func (s *PositionService) UpdatePosition(ctx context.Context, req *v12.UpdatePositionRequest) (*v12.Position, error) {
	return s.uc.Update(ctx, req)
}

func (s *PositionService) DeletePosition(ctx context.Context, req *v12.DeletePositionRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
