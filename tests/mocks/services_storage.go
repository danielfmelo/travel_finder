package mocks

import (
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/stretchr/testify/mock"
)

type ServiceStorageMock struct {
	mock.Mock
}

func (ssm *ServiceStorageMock) Save(record entity.Record) error {
	args := ssm.Called(record)
	return args.Error(0)
}
