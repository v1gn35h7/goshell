package respository

import (
	"time"

	"github.com/go-logr/zerologr"
	"github.com/kristoiv/gocqltable/recipes"
	"github.com/v1gn35h7/goshell/pkg/goshell"
)

type assetsRepository struct {
	logger      zerologr.Logger
	assetsTable struct {
		recipes.CRUD
	}
}

func AssetsRepository(logr zerologr.Logger) *assetsRepository {
	eTable := struct {
		recipes.CRUD // If you looked at the base example first, notice we replaced this line with the recipe
	}{
		recipes.CRUD{ // Here we didn't replace, but rather wrapped the table object in our recipe, effectively adding more methods to the end API
			TableInterface: Base().NewTable(
				"assets",        // The table name
				[]string{},      // Row keys
				nil,             // Range keys
				goshell.Asset{}, // We pass an instance of the asset struct that will be used as a type template during fetches.
			),
		},
	}

	return &assetsRepository{
		assetsTable: eTable,
		logger:      logr,
	}

}

func (r *assetsRepository) Update(asset goshell.Asset) (bool, error) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	asset.Synctime = time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST")

	_, err := r.assetsTable.Get(asset.Agentid)
	if err != nil {
		err := r.assetsTable.Insert(asset)
		if err != nil {
			r.logger.Error(err, "Failed to insert asset ")
			return false, err
		}
	} else {
		err := r.assetsTable.Update(asset)

		if err != nil {
			r.logger.Error(err, "Failed to update asset ")
			return false, err
		}
	}

	r.logger.Info("asset updated to cass")
	return true, nil
}

func (r *assetsRepository) List(hostId string) ([]*goshell.Asset, error) {
	defer func() {
		if ok := recover(); ok != nil {
			r.logger.Info("System Paniced", "rec:", ok)
		}
	}()
	list, err := r.assetsTable.List()

	if err != nil {
		r.logger.Error(err, "Failed to get assets")
		return nil, err
	}

	assets := list.([]*goshell.Asset)

	r.logger.Info("Assets fetched from cass")

	return assets, nil
}
