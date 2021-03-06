package header

import (
	"fmt"
	"strings"
)

type Require struct {
	Options []string
}

func (r *Require) String() string {
	return fmt.Sprintf("%s: %s", r.Name(), r.Value())
}

func (r *Require) Name() string {
	return "Require"
}

func (r *Require) Value() string {
	return strings.Join(r.Options, ", ")
}

func (r *Require) Equals(other interface{}) bool {
	if h, ok := other.(*Require); ok {
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

func (r *Require) Clone() Header {
	if r == nil {
		var newRequire *Require
		return newRequire
	}

	dup := make([]string, len(r.Options))
	copy(dup, r.Options)
	return &Require{dup}
}
