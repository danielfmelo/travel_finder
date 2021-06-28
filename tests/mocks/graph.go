package mocks

import (
	"strconv"

	"github.com/danielfmelo/travel_finder/entity"
)

type GraphMock struct {
	Path       []string
	Value      int
	AllRecords []entity.Record
}

func (g *GraphMock) AddEdge(origin, destiny string, value int) {
	g.AllRecords = append(g.AllRecords, entity.Record{origin, destiny, strconv.Itoa(value)})
	return
}

func (g *GraphMock) GetPath(origin, destiny string) ([]string, int) {
	return g.Path, g.Value
}
