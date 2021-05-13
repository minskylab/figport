package db

import (
	"github.com/asdine/storm/v3"
	"github.com/pkg/errors"
)

func openBoltDB(filepath string) (*storm.DB, error) {
	db, err := storm.Open(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
