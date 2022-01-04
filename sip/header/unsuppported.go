package header

import (
	"fmt"
	"strings"
)

type Unsupported struct {
	Options []string
}

func (r *Unsupported) String() string {
	return fmt.Sprintf("%s: %s", r.Name(), r.Value())
}

func (r *Unsupported) Name() string {
	return "Unsupported"
}

func (r *Unsupported) Value() string {
	return strings.Join(r.Options, ", ")
}

func (r *Unsupported) Equals(other interface{}) bool {
	if h, ok := other.(*Unsupported); ok {
		if r == h {
			return true
		}
		if h == nil && r != nil || h != nil && r == nil {
			return false
		}

		if len(r.Options) != len(h.Options) {
			return false
		}

		for i, opt := range r.Options {
			if opt != h.Options[i] {
				return false
			}
		}

		return true
	}
	return false
}

func (r *Unsupported) Clone() Header {
	if r == nil {
		var newUnsupported *Unsupported
		return newUnsupported
	}

	dup := make([]string, len(r.Options))
	copy(dup, r.Options)
	return &Unsupported{dup}
}
