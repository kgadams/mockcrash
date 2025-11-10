package mockcrash

//go:generate go tool mockery

type Value struct {
	Parameter string
}

type WithPtr interface {
	Do(*Value) error
}
