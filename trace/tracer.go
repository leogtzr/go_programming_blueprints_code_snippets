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

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New ...
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
