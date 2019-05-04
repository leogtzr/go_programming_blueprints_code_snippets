package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events thoughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type silentTracer struct{}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (silent *silentTracer) Trace(a ...interface{}) {}

// New ...
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Off ...
func Off() Tracer {
	return &silentTracer{}
}
