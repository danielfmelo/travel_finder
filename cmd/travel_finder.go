package main

import (
	"fmt"
	"os"

	"github.com/danielfmelo/travel_finder/deliveries/api"
	"github.com/danielfmelo/travel_finder/deliveries/command"
	"github.com/danielfmelo/travel_finder/deliveries/listener"
	"github.com/danielfmelo/travel_finder/deliveries/publisher"
	"github.com/danielfmelo/travel_finder/entity"
	"github.com/danielfmelo/travel_finder/repositories/datafile"
	"github.com/danielfmelo/travel_finder/repositories/graph"
	"github.com/danielfmelo/travel_finder/services"
)

func main() {
	if len(os.Args) < 2 {
		panic("not found initial database")
	}
	datafilePath := os.Args[1]
	queueBufferSize := 100
	writerQueue := make(chan entity.Record, queueBufferSize)

	d := datafile.New(datafilePath)
	allRecords, err := d.LoadAll()
	if err != nil {
		panic(fmt.Sprintf("error to load from database, error: %v", err))
	}
	g := graph.NewGraph()
	pub := publisher.New(writerQueue)
	finder := services.NewFinder(g, pub)
	err = finder.SaveAll(allRecords)
	if err != nil {
		panic(fmt.Sprintf("error to save all records to storage, error: %v", err))
	}

	c := command.New(finder)
	s := services.NewStorage(d)
	listenerSub := listener.NewReader(s)
	listenerSub.Start(writerQueue)
	a := api.New(finder)
	go c.Start()
	a.Start()
}
