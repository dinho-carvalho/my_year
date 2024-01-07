package savefile

import (
	"context"
	"io"

	"github.com/dinho-carvalho/my_year/internal/domain/interfaces"
)

type (
	UseCaseImpl struct {
		Gateway interfaces.StorageGateway
	}

	UseCase interface {
		Execute(ctx context.Context, file io.Reader, filename string) error
	}
)

func NewUseCase(gateway interfaces.StorageGateway) *UseCaseImpl {
	return &UseCaseImpl{Gateway: gateway}
}

func (u *UseCaseImpl) Execute(ctx context.Context, file io.Reader, filename string) error {
	return u.Gateway.SaveFile(ctx, file, filename)
}
