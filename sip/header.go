package sip

type Header interface {
	String() string
	Name() string
	Value() string
	Equals() bool
	Clone() Header
}
