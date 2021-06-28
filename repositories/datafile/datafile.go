package datafile

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/danielfmelo/travel_finder/entity"
)

type Database struct {
	filePath string
}

func (d *Database) LoadAll() ([]entity.Record, error) {
	csvfile, err := os.Open(d.filePath)
	if err != nil {
		return nil, err
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)

	rAll, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	records := []entity.Record{}
	for _, rec := range rAll {
		records = append(records, entity.Record{rec[0], rec[1], rec[2]})
	}
	return records, nil
}

func (d *Database) Save(record entity.Record) error {
	r := []string{record.Origin, record.Destination, record.Value}
	csvfile, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer csvfile.Close()
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	w := csv.NewWriter(csvfile)
	if err := w.Write(convertRecordToUpperCase(r)); err != nil {
		return err
	}
	w.Flush()
	return nil
}

func convertRecordToUpperCase(record []string) []string {
	var upper []string
	for _, r := range record {
		upper = append(upper, strings.ToUpper(r))
	}
	return upper
}

func New(databasePath string) *Database {
	return &Database{
		filePath: databasePath,
	}
}
