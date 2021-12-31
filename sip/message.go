package sip

import "github.com/gozh-io/sip/sip/header"

type Message interface {
	// Returns a slice of all headers of the given type.
	// If there are no headers of the requested type, returns an empty slice.

	GetHeader(string) []header.Header

	// Adds a header to this message.

	SetHeader(header.Header)

	// Return all headers attached to the message, as a slice.

	AllHeaders() []header.Header

	// Remove the specified header from the message.

	RemoveHeader(header header.Header) error

	// Yields a flat, string representation of the SIP message suitable for sending out over the wire.

	String() string

	// Get the body of the message, as a string.

	GetBody() string

	// Set the body of the message.

	SetBody(body string)
}
