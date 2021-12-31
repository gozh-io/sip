package header

import "fmt"

type CallID string

func (c *CallID) String() string {
	return fmt.Sprintf("%s: %s", c.Name(), c.Value())
}

func (c *CallID) Name() string {
	return "Call-ID"
}

func (c CallID) Value() string {
	return string(c)
}

func (c *CallID) Equals(other interface{}) bool {
	if h, ok := other.(CallID); ok {
		if c == nil {
			return false
		}
		return *c == h
	}
	if h, ok := other.(*CallID); ok {
		if c == h {
			return true
		}
		if c == nil && h != nil || c != nil && h == nil {
			return false
		}

		return *c == *h
	}
	return false
}

func (c *CallID) Clone() Header {
	return c
}
