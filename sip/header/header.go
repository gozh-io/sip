package header

type Header interface {
	String() string
	Name() string
	Value() string
	Equals(other interface{}) bool
	Clone() Header
}
