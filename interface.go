package mockcrash

import "fmt"

//go:generate go tool mockery

type Value struct {
	Parameter string
}

func (v Value) String() string {
	return fmt.Sprintf("Value(%s)", v.Parameter)
}

type WithPtr interface {
	Do(*Value) (Value, error)
}
