package article

import "context"

var insSinglePage = sSinglePage{}

func SinglePage() *sSinglePage {
	return &insSinglePage
}

type sSinglePage struct {
}

func (s *sSinglePage) CreateCategory(ctx context.Context) {

}
