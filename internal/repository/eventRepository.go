package respository

import (
	"github.com/go-logr/zerologr"
	"github.com/goShell/pkg/goshell"
	"github.com/kristoiv/gocqltable/recipes"
)

type eventRepository struct {
	logger     zerologr.Logger
	eventTable struct {
		recipes.CRUD
	}
}

func EventRepository(logr zerologr.Logger) *eventRepository {
	eTable := struct {
		recipes.CRUD // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			TableInterface: Base().NewTable(
				"trooper_events", // The table name
				[]string{"id"},   // Row keys
				nil,              // Range keys
				goshell.Events{}, // We pass an instance of the user struct that will be used as a type template during fetches.
			),
		},
	}

	return &eventRepository{
		eventTable: eTable,
		logger:     logr,
	}
}

func (rep *eventRepository) AddEvents(event goshell.Events) {
	defer func() {
		if ok := recover(); ok != nil {
			rep.logger.Info("System Paniced", "rec:", ok)
		}
	}() 
	err := rep.eventTable.Insert(event)

	if err != nil {
		rep.logger.Error(err, "Failed to insert event ")
	}

	rep.logger.Info("Event added to cass")
	panic("test")
}
