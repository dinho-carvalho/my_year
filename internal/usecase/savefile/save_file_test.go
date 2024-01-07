package savefile

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
}

func (s *StorageMock) SaveFile(ctx context.Context, file io.Reader, filename string) error {
	args := s.Called(ctx, file, filename)
	return args.Error(0)
}

type SuiteSaveFileUseCase struct {
	suite.Suite
	SuiteSaveFileUseCase *UseCaseImpl
	StorageMock          *StorageMock
}

func (t *SuiteSaveFileUseCase) SetupTest() {
	storageMock := new(StorageMock)

	t.SuiteSaveFileUseCase = NewUseCase(storageMock)
	t.StorageMock = storageMock
}

func TestSuiteSaveFileUseCase(t *testing.T) {
	suite.Run(t, new(SuiteSaveFileUseCase))
}

func (s *SuiteSaveFileUseCase) clearMocks() {
	s.StorageMock.ExpectedCalls = nil
}

func (s *SuiteSaveFileUseCase) TestSuiteSaveFileUseCase() {
	s.Run("Should handle save file", func() {
		defer s.clearMocks()

		ctx := context.Background()

		s.StorageMock.On("SaveFile", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := s.SuiteSaveFileUseCase.Execute(ctx, nil, "")

		s.NoError(err)
	})

	s.Run("Should handle save file error", func() {
		defer s.clearMocks()

		ctx := context.Background()

		s.StorageMock.On("SaveFile", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("some error"))

		err := s.SuiteSaveFileUseCase.Execute(ctx, nil, "")

		s.ErrorContains(err, "some error")
	})
}
