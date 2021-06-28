package mocks

import (
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/stretchr/testify/mock"
)

type PublisherMock struct {
	mock.Mock
	Record entity.Record
}

func (pm *PublisherMock) Send(record entity.Record) {
	pm.Called(record)
	pm.Record = record
}
