package repositories

import "github.com/danielfmelo/travel_finder/entity"

type Graph interface {
	AddEdge(origin, destiny string, value int)
	GetPath(origin, destiny string) ([]string, int)
}

type Storage interface {
	LoadAll() ([]entity.Record, error)
	Save(record entity.Record) error
}
