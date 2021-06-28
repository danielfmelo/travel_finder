package services_test

import (
	"testing"

	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/services"
	"github.com/danielfmelo/travel_finder/tests/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	repoMock := &mocks.RepoStorageMock{}
	s := services.NewStorage(repoMock)
	record := entity.Record{"test", "mock", "123"}
	repoMock.On("Save", record).Return(nil).Once()
	err := s.Save(record)
	assert.NoError(t, err)
}

func TestSaveShouldReturnErrorOnRepo(t *testing.T) {
	repoMock := &mocks.RepoStorageMock{}
	s := services.NewStorage(repoMock)
	record := entity.Record{"test", "mock", "123"}
	fakeErr := errors.New("some error")
	repoMock.On("Save", record).Return(fakeErr).Once()
	err := s.Save(record)
	assert.Equal(t, fakeErr, err)
}
