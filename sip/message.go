package sip

import (
	"github.com/gozh-io/sip/sip/header"
	"github.com/gozh-io/sip/sip/primitive"
)

type Message interface {
	MessageID() primitive.MessageID

	// StartLine start line
	StartLine() string

	// GetHeader  Returns a slice of all headers of the given type.
	// GetHeader  If there are no headers of the requested type, returns an empty slice.
	GetHeader(string) []header.Header

	// SetHeader  Adds a header to this message.
	SetHeader(header.Header)

	// AllHeaders Return all headers attached to the message, as a slice.
	AllHeaders() []header.Header

	// RemoveHeader Remove the specified header from the message.
	RemoveHeader(header header.Header) error

	// String Yields a flat, string representation of the SIP message suitable for sending out over the wire.
	String() string

	// GetBody Get the body of the message, as a string.
	GetBody() string

	// SetBody Set the body of the message.
	SetBody(body string)
}

type message struct {
	headers []header.Header
}
