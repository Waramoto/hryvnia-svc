package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/Waramoto/hryvnia-svc/internal/data"
)

type db struct {
	raw *pgdb.DB
}

func NewDB(rawDB *pgdb.DB) data.DB {
	return &db{
		raw: rawDB,
	}
}

func (db *db) New() data.DB {
	return NewDB(db.raw.Clone())
}

func (db *db) Subscribers() data.SubscribersQ {
	return NewSubscribersQ(db.raw)
}

func (db *db) Transaction(fn func() error) error {
	return db.raw.Transaction(func() error {
		return fn()
	})
}
