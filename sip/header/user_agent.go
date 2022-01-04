package header

import "fmt"

type UserAgent string

func (m *UserAgent) String() string {
	return fmt.Sprintf("%s: %s", m.Name(), m.Value())
}

func (m *UserAgent) Name() string {
	return "User-Agent"
}

func (m UserAgent) Value() string {
	return string(m)
}

func (m *UserAgent) Equals(other interface{}) bool {
	if h, ok := other.(UserAgent); ok {
		if m == nil {
			return false
		}

		return *m == h
	}
	if h, ok := other.(*UserAgent); ok {
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

func (m *UserAgent) Clone() Header {
	return m
}

