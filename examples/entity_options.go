package examples

import (
	"fmt"
)

type Entity struct {
	name string
	opt  string
}

func (e *Entity) Show() {
	fmt.Printf("Name:%s,Opt:%s", e.name, e.opt)
}

// use optional option list param
func NewMyEntity(name string, options ...Option) *Entity {
	e := &Entity{ // default entity if no options provided
		name: name,
		opt:  "default opt",
	}

	// apply option list
	for _, o := range options {
		o.Apply(e)
	}

	return e
}

type Option interface {
	Apply(*Entity)
}

type OptionFunc func(*Entity)

// OptionFunc implements Apply, it can be use as Option
// a function can have methods, quite tricky.
func (f OptionFunc) Apply(e *Entity) {
	f(e)
}

func WithCustomOpt(opt string) Option {
	// OptionFunc is a type, OptionFunc(xx) converts xx into its type.
	return OptionFunc(func(e *Entity) {
		e.opt = opt
	})
}

func ExampleEntityOption() {
	e := NewMyEntity("foo", WithCustomOpt("hi"))
	e.Show()
}
