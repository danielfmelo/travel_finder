package datafile_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/repositories/datafile"
	"github.com/stretchr/testify/assert"
)

func TestLoadAll(t *testing.T) {
	testFile := "test.csv"
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatal("error to create testFile")
	}
	defer func(testFile string) {
		if err = os.Remove(testFile); err != nil {
			t.Fatal("error to remove testFile")
		}
	}(testFile)
	rawRecord := []string{"GRU", "FLN", "40"}
	expectedRecord := entity.Record{"GRU", "FLN", "40"}
	writer := csv.NewWriter(file)
	if err = writer.Write(rawRecord); err != nil {
		t.Fatal("error to write testFile")
	}
	writer.Flush()
	file.Close()

	d := datafile.New(testFile)
	records, err := d.LoadAll()
	if err != nil {
		t.Errorf("err should be nil, but it is: %v", err)
	}
	assert.Equal(t, expectedRecord, records[0])
}

func TestWrite(t *testing.T) {
	testFile := "test.csv"
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatal("error to create testFile")
	}
	defer func(testFile string) {
		if err = os.Remove(testFile); err != nil {
			t.Fatal("error to remove testFile")
		}
	}(testFile)
	file.Close()
	expectedRecord := entity.Record{"GRU", "FLN", "40"}
	d := datafile.New(testFile)
	err = d.Save(expectedRecord)
	if err != nil {
		t.Fatal("error to create testFile")
	}
	records, err := d.LoadAll()
	if err != nil {
		t.Errorf("err should be nil, but it is: %v", err)
	}
	assert.Equal(t, expectedRecord, records[0])
}
