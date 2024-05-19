package data

import (
	"time"

	"github.com/Waramoto/hryvnia-svc/internal/types"
)

type SubscribersQ interface {
	New() SubscribersQ

	Insert(subscriber Subscriber) error

	Select() ([]Subscriber, error)

	UpdateLastSend(lastSend time.Time) error
	UpdateStatus(status types.Status) error

	FilterByEmails(emails ...string) SubscribersQ
}

type Subscriber struct {
	ID       int64        `db:"id" structs:"-"`
	Email    string       `db:"email" structs:"email"`
	LastSend time.Time    `db:"last_send" structs:"last_send"`
	Status   types.Status `db:"status" structs:"status"`
}
