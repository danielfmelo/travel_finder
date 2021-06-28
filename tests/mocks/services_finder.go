package mocks

import (
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/stretchr/testify/mock"
)

type FinderMock struct {
	mock.Mock
}

func (sm *FinderMock) Save(record entity.Record) error {
	args := sm.Called(record)
	return args.Error(0)
}

func (sm *FinderMock) GetSmallestPriceAndRoute(origin, destin string) (entity.CheapestRoute, error) {
	args := sm.Called(origin, destin)
	return args.Get(0).(entity.CheapestRoute), args.Error(1)
}

func (sm *FinderMock) SaveAll(records []entity.Record) error {
	args := sm.Called(records)
	return args.Error(0)
}
