package header

import "fmt"

type MaxForwards uint32

func (m *MaxForwards) String() string {
	return fmt.Sprintf("%s: %s", m.Name(), m.Value())
}

func (m *MaxForwards) Name() string {
	return "Max-Forwards"
}

func (m MaxForwards) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m *MaxForwards) Equals(other interface{}) bool {
	if h, ok := other.(MaxForwards); ok {
		if m == nil {
			return false
		}

		return *m == h
	}
	if h, ok := other.(*MaxForwards); ok {
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

func (m *MaxForwards) Clone() Header {
	return m
}
