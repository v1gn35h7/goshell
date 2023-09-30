package respository

import (
	"github.com/go-logr/zerologr"
	"github.com/kristoiv/gocqltable/recipes"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

type scriptsRepository struct {
	logger       zerologr.Logger
	scriptsTable struct {
		recipes.CRUD
	}
}

func ScriptsRepository(logr zerologr.Logger) *scriptsRepository {
	eTable := struct {
		recipes.CRUD // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			TableInterface: Base().NewTable(
				"scripts",        // The table name
				[]string{"id"},   // Row keys
				nil,              // Range keys
				goshell.Script{}, // We pass an instance of the user struct that will be used as a type template during fetches.
			),
		},
	}

	return &scriptsRepository{
		scriptsTable: eTable,
		logger:       logr,
	}

}

func (r *scriptsRepository) Save(script goshell.Script) (bool, error) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	err := r.scriptsTable.Insert(script)

	if err != nil {
		r.logger.Error(err, "Failed to insert script ")
		return false, err
	}

	r.logger.Info("Script added to cass")

	return true, nil
}

func (r *scriptsRepository) List(hostId string) ([]*goshell.Script, error) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	list, err := r.scriptsTable.List()

	if err != nil {
		r.logger.Error(err, "Failed to insert script ")
		return nil, err
	}

	scripts := list.([]*goshell.Script)

	r.logger.Info("Script added to cass")

	return scripts, nil
}
