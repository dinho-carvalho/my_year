package interfaces

import (
	"context"
	"io"
)

type StorageGateway interface {
	SaveFile(ctx context.Context, file io.Reader, filename string) error
}
