package data

//go:generate mockery --case=underscore --name=DB
type DB interface {
	New() DB

	Subscribers() SubscribersQ

	Transaction(func() error) error
}
