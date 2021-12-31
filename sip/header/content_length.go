package header

import "fmt"

type ContentLength uint32

func (m *ContentLength) String() string {
	return fmt.Sprintf("%s: %s", m.Name(), m.Value())
}

func (m *ContentLength) Name() string {
	return "Content-Length"
}

func (m ContentLength) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m *ContentLength) Equals(other interface{}) bool {
	if h, ok := other.(ContentLength); ok {
		if m == nil {
			return false
		}

		return *m == h
	}
	if h, ok := other.(*ContentLength); ok {
		if m == h {
			return true
		}
		if m == nil && h != nil || m != nil && h == nil {
			return false
		}

		return *m == *h
	}

	return false
}

func (m *ContentLength) Clone() Header {
	return m
}

