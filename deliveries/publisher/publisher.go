package publisher

import "github.com/danielfmelo/travel_finder/entity"

type Record interface {
	Send(record entity.Record)
}

type record struct {
	queue chan entity.Record
}

func New(queue chan entity.Record) Record {
	return &record{queue: queue}
}

func (r *record) Send(record entity.Record) {
	r.queue <- record
}
