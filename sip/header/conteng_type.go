package header

import "fmt"

type ContentType string

func (ct *ContentType) String() string { return fmt.Sprintf("%s: %s", ct.Name(), ct.Value()) }

func (ct *ContentType) Name() string { return "Content-Type" }

func (ct ContentType) Value() string { return string(ct) }

func (ct *ContentType) Clone() Header { return ct }

func (ct *ContentType) Equals(other interface{}) bool {
	if h, ok := other.(ContentType); ok {
		if ct == nil {
			return false
		}

		return *ct == h
	}
	if h, ok := other.(*ContentType); ok {
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

