package article

import (
	"context"
	"ice_flame/internal/service"
)

var insSinglePage = sSinglePage{}

func SinglePage() *sSinglePage {
	return &insSinglePage
}

func init() {
	service.RegisterSinglePage(SinglePage())
}

type sSinglePage struct {
}

func (s *sSinglePage) CreateCategory(ctx context.Context) {
}

func (s *sSinglePage) GetCategoryById(ctx context.Context) {
}

func (s *sSinglePage) ModifyCategoryById(ctx context.Context) {
}

func (s *sSinglePage) ListCategoryById(ctx context.Context) {
}

func (s *sSinglePage) DeleteCategoryById(ctx context.Context) {
}
