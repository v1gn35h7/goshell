package respository

import (
	"fmt"

	"github.com/go-logr/zerologr"
	"github.com/kristoiv/gocqltable/recipes"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

type resultsRepository struct {
	logger       zerologr.Logger
	resultsTable struct {
		recipes.CRUD
	}
}

func ResultsRepository(logr zerologr.Logger) *resultsRepository {
	eTable := struct {
		recipes.CRUD // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			TableInterface: Base().NewTable(
				"results",        // The table name
				[]string{"id"},   // Row keys
				nil,              // Range keys
				goshell.Output{}, // We pass an instance of the asset struct that will be used as a type template during fetches.
			),
		},
	}

	return &resultsRepository{
		resultsTable: eTable,
		logger:       logr,
	}

}

func (r *resultsRepository) Save(output goshell.Output) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	err := r.resultsTable.Insert(output)

	if err != nil {
		r.logger.Error(err, "Failed to insert output")
	}

	r.logger.Info("Output added to cass")
}

func (r *resultsRepository) Find(query string) ([]*goshell.Output, error) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	list, err := r.resultsTable.List()

	if err != nil {
		r.logger.Error(err, "Failed to search results...")
		return nil, err
	}

	results := list.([]*goshell.Output)

	msg := fmt.Sprintf("Found %d results", len(results))

	r.logger.Info(msg)

	return results, nil
}
