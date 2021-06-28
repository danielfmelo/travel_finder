package services

import (
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/repositories"
)

type storage struct {
	datafileStorage repositories.Storage
}

func (s *storage) Save(record entity.Record) error {
	return s.datafileStorage.Save(record)
}
