package header

import (
	"fmt"
	"strings"
)

type Supported struct {
	Options []string
}

func (r *Supported) String() string {
	return fmt.Sprintf("%s: %s", r.Name(), r.Value())
}

func (r *Supported) Name() string {
	return "Supported"
}

func (r *Supported) Value() string {
	return strings.Join(r.Options, ", ")
}

func (r *Supported) Equals(other interface{}) bool {
	if h, ok := other.(*Supported); ok {
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

func (r *Supported) Clone() Header {
	if r == nil {
		var newSupported *Supported
		return newSupported
	}

	dup := make([]string, len(r.Options))
	copy(dup, r.Options)
	return &Supported{dup}
}
