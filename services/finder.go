package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielfmelo/travel_finder/deliveries/publisher"
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/repositories"
	"github.com/pkg/errors"
)

type finder struct {
	graph     repositories.Graph
	recordPub publisher.Record
}

func (f *finder) SaveAll(records []entity.Record) error {
	for _, record := range records {
		if err := validateAndAdjustRecord(&record); err != nil {
			return errors.Wrap(err, "error to extract data from record")
		}
		if err := f.saveOne(record); err != nil {
			return errors.Wrap(err, "error to save all records")
		}
	}
	return nil
}

func (f *finder) saveOne(record entity.Record) error {
	val, err := convertStrToInt(record.Value)
	if err != nil {
		return err
	}
	f.graph.AddEdge(record.Origin, record.Destination, val)
	return nil
}

func (f *finder) Save(record entity.Record) error {
	if err := validateAndAdjustRecord(&record); err != nil {
		return errors.Wrap(err, "error to extract data from record")
	}
	val, err := convertStrToInt(record.Value)
	if err != nil {
		return err
	}
	f.graph.AddEdge(record.Origin, record.Destination, val)
	//f.recordPub.Send(record)
	return nil
}

func (f *finder) GetSmallestPriceAndRoute(origin, destination string) (entity.CheapestRoute, error) {
	ori := strings.ToUpper(origin)
	dest := strings.ToUpper(destination)
	locals, value := f.graph.GetPath(ori, dest)
	if locals == nil {
		return entity.CheapestRoute{}, entity.ErrNotFound
	}
	cr := entity.CheapestRoute{
		Path:  fmt.Sprintf("%s", strings.Join(locals, " - ")),
		Value: value,
	}
	return cr, nil
}

func validateAndAdjustRecord(record *entity.Record) error {
	if record.Origin == "" || record.Destination == "" {
		return errors.New("origin and destination shouln'd be empty")
	}
	record.Origin = strings.ToUpper(record.Origin)
	record.Destination = strings.ToUpper(record.Destination)
	if _, err := convertStrToInt(record.Value); err != nil {
		return err
	}
	return nil
}

func convertStrToInt(value string) (int, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return -1, errors.New("error to extrac value")
	}
	return val, err
}
