package sip

type Transport interface {
	Message() <-chan Message
	Send(msg Message) error
	IsReliable(network string) bool
	IsStreamed(network string) bool
}
