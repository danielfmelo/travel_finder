package services_test

import (
	"fmt"
	"testing"

	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/services"
	"github.com/danielfmelo/travel_finder/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSaveAllShouldWork(t *testing.T) {
	expectedRecords := []entity.Record{
		{"FLN", "GRU", "40"},
	}
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	s := services.NewFinder(g, p)
	p.On("Send", expectedRecords[0]).Return().Once()
	err := s.SaveAll(expectedRecords)
	if err != nil {
		t.Errorf("error should be nil but it is: %v", err)
	}
	assert.Equal(t, expectedRecords, g.AllRecords)
}

func TestSaveAllShouldReturnError(t *testing.T) {
	expectedRecords := []entity.Record{
		{"FLN", "GRU", "dasds"},
	}
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	s := services.NewFinder(g, p)
	p.On("Send", expectedRecords[0]).Return().Once()
	err := s.SaveAll(expectedRecords)
	expectedErrorMsg := "error to extract data from record: error to extrac value"
	if err.Error() != expectedErrorMsg {
		t.Errorf("err should be %s, but got %s", expectedErrorMsg, err.Error())
	}
}

func TestSaveRecord(t *testing.T) {
	ori := "FLN"
	dest := "GRU"
	value := "40"
	expectedRecords := []entity.Record{
		{ori, dest, value},
	}
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	s := services.NewFinder(g, p)
	record := entity.Record{ori, dest, value}
	p.On("Send", record).Return().Once()
	err := s.Save(record)
	if err != nil {
		t.Errorf("error should be nil but it is: %v", err)
	}
	assert.Equal(t, expectedRecords, g.AllRecords)
}

func TestSaveRecordShouldReturnError(t *testing.T) {
	ori := "FLN"
	dest := "GRU"
	value := "wrong"
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	s := services.NewFinder(g, p)
	record := entity.Record{ori, dest, value}
	p.On("Send", record).Return().Once()
	err := s.Save(record)
	expectedErrorMsg := "error to extract data from record: error to extrac value"
	if err.Error() != expectedErrorMsg {
		t.Errorf("err should be %s, but got %s", expectedErrorMsg, err.Error())
	}
}

func TestGetSmallestPriceAndRoute(t *testing.T) {
	ori := "GRU"
	dest := "FLN"
	path := []string{ori, dest}
	expectedValue := 40
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	g.Path = path
	g.Value = expectedValue
	s := services.NewFinder(g, p)
	cheapestRoute, err := s.GetSmallestPriceAndRoute(ori, dest)
	assert.NoError(t, err)
	expectedPath := fmt.Sprintf("GRU - FLN")
	assert.Equal(t, expectedPath, cheapestRoute.Path)
	assert.Equal(t, expectedValue, cheapestRoute.Value)
}

func TestGetSmallestPriceAndRouteShouldReturnNotFound(t *testing.T) {
	ori := "GRU"
	dest := "FLN"
	g := &mocks.GraphMock{}
	p := &mocks.PublisherMock{}
	g.Path = nil
	s := services.NewFinder(g, p)
	_, err := s.GetSmallestPriceAndRoute(ori, dest)
	assert.Equal(t, entity.ErrNotFound, err)
}
