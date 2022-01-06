package sip

import (
	"bytes"
	"strings"
	"sync"

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
	*headers
	mu         sync.RWMutex
	messageID  primitive.MessageID
	sipVersion string
	body       string
	startLine  func() string
	tp         string
	src        string
	dest       string
}

func (a *message) MessageID() primitive.MessageID {
	return a.messageID
}

func (a *message) StartLine() string {
	return a.startLine()
}

func (a *message) String() string {
	var buff bytes.Buffer

	buff.WriteString(a.StartLine() + "\r\n")

	a.mu.RLock()
	buff.WriteString(a.headers.String())
	a.mu.RUnlock()

	buff.WriteString("\r\n" + a.Body())

	return buff.String()
}

func (a *message) Body() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.body
}

func (a *message) SipVersion() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.sipVersion
}

func (a *message) SetSipVersion(version string) {
	a.mu.Lock()
	a.sipVersion = version
	a.mu.Unlock()
}

// SetBody sets message body, calculates it length and add 'Content-Length' header.
func (a *message) SetBody(body string, setContentLength bool) {
	a.mu.Lock()
	a.body = body
	a.mu.Unlock()
	if setContentLength {
		hdrs := a.GetHeaders("Content-Length")
		if len(hdrs) == 0 {
			length := header.ContentLength(len(body))
			a.Append(&length)
		} else {
			length := header.ContentLength(len(body))
			a.ReplaceHeader("Content-Length", []header.Header{&length})
		}
	}
}

func (a *message) Transport() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.tp
}

func (a *message) SetTransport(tp string) {
	a.mu.Lock()
	a.tp = strings.ToUpper(tp)
	a.mu.Unlock()
}

func (a *message) Source() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.src
}

func (a *message) SetSource(src string) {
	a.mu.Lock()
	a.src = src
	a.mu.Unlock()
}

func (a *message) Destination() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.dest
}

func (a *message) SetDestination(dest string) {
	a.mu.Lock()
	a.dest = dest
	a.mu.Unlock()
}

type MessageMapper func(msg Message)