package mocks

import "github.com/danielfmelo/travel_finder/entity"

type FinderMock struct {
	FakeErr       error
	CheapestRoute entity.CheapestRoute
	Value         int
	Records       []entity.Record
}

func (sm *FinderMock) Save(entity.Record) error {
	return sm.FakeErr
}

func (sm *FinderMock) GetSmallestPriceAndRoute(origin, destin string) (entity.CheapestRoute, error) {
	return sm.CheapestRoute, sm.FakeErr
}

func (sm *FinderMock) SaveAll(records []entity.Record) error {
	sm.Records = records
	return sm.FakeErr
}
