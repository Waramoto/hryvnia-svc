package types

type Status int8

const (
	StatusNotSent Status = iota
	StatusSent
)

func (s Status) ToInt8() int8 {
	return int8(s)
}
