package db

import (
	"github.com/asdine/storm/v3"
	"github.com/pkg/errors"
)

type Repository struct {
	filepath string
	db       *storm.DB
}

type RepositoryOption func(*Repository)

func WithFilepath(filepath string) RepositoryOption {
	return func(r *Repository) {
		r.filepath = filepath
	}
}

func NewRepository(options ...RepositoryOption) (*Repository, error) {
	r := &Repository{}

	for _, opt := range options {
		opt(r)
	}

	db, err := openBoltDB(r.filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r.db = db

	return r, nil
}
