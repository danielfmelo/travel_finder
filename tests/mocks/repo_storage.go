package mocks

import (
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/stretchr/testify/mock"
)

type RepoStorageMock struct {
	mock.Mock
}

func (rsm *RepoStorageMock) LoadAll() ([]entity.Record, error) {
	args := rsm.Called()
	return args.Get(0).([]entity.Record), args.Error(1)
}

func (rsm *RepoStorageMock) Save(record entity.Record) error {
	args := rsm.Called(record)
	return args.Error(0)
}
