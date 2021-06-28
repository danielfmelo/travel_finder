package listener

import (
	"fmt"

	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/services"
)

type Reader struct {
	storageService services.Storage
}

func NewReader(storageService services.Storage) *Reader {
	return &Reader{
		storageService: storageService,
	}
}

func (r *Reader) Start(writeChn chan entity.Record) {
	go func() {
		for {
			select {
			case record := <-writeChn:
				if err := r.storageService.Save(record); err != nil {
					fmt.Printf("error saving record on storage: %v\n", err)
				}
			}
		}
	}()
}
