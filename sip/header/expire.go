package header

import "fmt"

type Expire uint32

func (m *Expire) String() string {
	return fmt.Sprintf("%s: %s", m.Name(), m.Value())
}

func (m *Expire) Name() string {
	return "Expire"
}

func (m Expire) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m *Expire) Equals(other interface{}) bool {
	if h, ok := other.(Expire); ok {
		if m == nil {
			return false
		}

		return *m == h
	}
	if h, ok := other.(*Expire); ok {
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

func (m *Expire) Clone() Header {
	return m
}

