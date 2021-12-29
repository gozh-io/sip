package sip

type Stringer interface {
	String() string
	Equals(other interface{}) bool
}

type String struct {
	Value string
}

type Port uint16

func NewPort(port uint16) *Port {
	p := Port(port)
	return &p
}

func (str *String) String() string {
	return str.Value
}

func (str *String) Equals(other interface{}) bool {
	if v, ok := other.(string); ok {
		return str.Value == v
	}
	if v, ok := other.(String); ok {
		return str.Value == v.Value
	}

	return false
}
