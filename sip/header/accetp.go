package header

import "fmt"

type Accept string

func (ct *Accept) String() string { return fmt.Sprintf("%s: %s", ct.Name(), ct.Value()) }

func (ct *Accept) Name() string { return "Accept" }

func (ct Accept) Value() string { return string(ct) }

func (ct *Accept) Clone() Header { return ct }

func (ct *Accept) Equals(other interface{}) bool {
	if h, ok := other.(Accept); ok {
		if ct == nil {
			return false
		}

		return *ct == h
	}
	if h, ok := other.(*Accept); ok {
		if ct == h {
			return true
		}
		if ct == nil && h != nil || ct != nil && h == nil {
			return false
		}

		return *ct == *h
	}

	return false
}

