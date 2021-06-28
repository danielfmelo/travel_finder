package services

import (
	"github.com/danielfmelo/travel_finder/deliveries/publisher"
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/repositories"
)

type Finder interface {
	SaveAll(records []entity.Record) error
	Save(record entity.Record) error
	GetSmallestPriceAndRoute(origin, destination string) (entity.CheapestRoute, error)
}

type Storage interface {
	Save(record entity.Record) error
}

func NewFinder(graph repositories.Graph, recordPub publisher.Record) Finder {
	return &finder{
		graph:     graph,
		recordPub: recordPub,
	}
}

func NewStorage(datafileStorage repositories.Storage) Storage {
	return &storage{
		datafileStorage: datafileStorage,
	}
}
