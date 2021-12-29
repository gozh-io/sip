package sip

type Message interface {
	// Returns a slice of all headers of the given type.
	// If there are no headers of the requested type, returns an empty slice.

	GetHeader(string) []Header

	// Adds a header to this message.

	SetHeader(Header)

	// Return all headers attached to the message, as a slice.

	AllHeaders() []Header

	// Remove the specified header from the message.

	RemoveHeader(header Header) error

	// Yields a flat, string representation of the SIP message suitable for sending out over the wire.

	String() string

	// Get the body of the message, as a string.

	GetBody() string

	// Set the body of the message.

	SetBody(body string)
}
