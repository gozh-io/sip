package header

import (
	"fmt"
	"github.com/bean-du/sip/pkg/sip"
	"github.com/gozh-io/sip/sip/primitive"
	"strings"
)

type Allow []primitive.RequestMethod

func (a Allow) Name() string {
	return "Allow"
}

func (a Allow) Value() string {
	parts := make([]string, 0)
	for _, p := range a {
		parts = append(parts, p.String())
	}

	return strings.Join(parts, ", ")
}

func (a Allow) Clone() sip.Header {
	if a == nil {
		var allow Allow
		return allow
	}

	newAllow := make(Allow, len(a))
	copy(newAllow, a)
	return newAllow
}

func (a Allow) String() string {
	return fmt.Sprintf("%s: %s", a.Name(), a.Value())
}

func (a Allow) Equals(other interface{}) bool {
	if h, ok := other.(Allow); ok {
		if len(a) != len(h) {
			return false
		}

		for i, v := range a {
			if h[i] != v {
				return false
			}
		}
		return true
	}
	return false
}
