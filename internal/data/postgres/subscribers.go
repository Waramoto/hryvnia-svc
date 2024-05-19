package data

import (
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/lib/pq"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/Waramoto/hryvnia-svc/internal/data"
	"github.com/Waramoto/hryvnia-svc/internal/types"
)

const (
	subscribersTable = "subscribers"

	subscribersID       = "id"
	subscribersEmail    = "email"
	subscribersLastSend = "last_send"
	subscribersStatus   = "status"

	ErrUniqueViolation = "unique_violation"
)

var (
	ErrAlreadyExists = errors.New("subscriber already exists")
)

type subscribersQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func NewSubscribersQ(db *pgdb.DB) data.SubscribersQ {
	return &subscribersQ{
		db:       db,
		selector: squirrel.Select("*").From(subscribersTable),
		updater:  squirrel.Update(subscribersTable),
	}
}

func (q *subscribersQ) New() data.SubscribersQ {
	return NewSubscribersQ(q.db.Clone())
}

func (q *subscribersQ) Insert(subscriber data.Subscriber) error {
	valuesMap := structs.Map(&subscriber)
	err := q.db.Exec(squirrel.Insert(subscribersTable).SetMap(valuesMap))
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code.Name() == ErrUniqueViolation {
		return ErrAlreadyExists
	}
	return err
}

func (q *subscribersQ) Select() ([]data.Subscriber, error) {
	var subscribers []data.Subscriber
	err := q.db.Select(&subscribers, q.selector)
	return subscribers, err
}

func (q *subscribersQ) UpdateLastSend(lastSend time.Time) error {
	return q.db.Exec(q.updater.Set(subscribersLastSend, lastSend))
}

func (q *subscribersQ) UpdateStatus(status types.Status) error {
	return q.db.Exec(q.updater.Set(subscribersStatus, status))
}

func (q *subscribersQ) FilterByEmails(emails ...string) data.SubscribersQ {
	q.selector = q.selector.Where(squirrel.Eq{subscribersEmail: emails})
	q.updater = q.updater.Where(squirrel.Eq{subscribersEmail: emails})
	return q
}
