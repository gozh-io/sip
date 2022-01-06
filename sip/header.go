package sip

import (
	"bytes"
	"strings"
	"sync"

	"github.com/gozh-io/sip/sip/header"
)

type headers struct {
	mu sync.Mutex
	// The logical SIP headers attached to this message.
	headers map[string][]header.Header
	// The order the headers should be display in
	order []string
}

func newHeaders(hds []header.Header) *headers {
	hs := new(headers)
	hs.headers = make(map[string][]header.Header)
	hs.order = make([]string, 0)
	for _, h := range hds {
		hs.Append(h)
	}
	return hs
}

func (a *headers) Append(h header.Header) {
	name := strings.ToLower(h.Name())
	a.mu.Lock()
	if _, ok := a.headers[name]; ok {
		a.headers[name] = append(a.headers[name], h)
	} else {
		a.headers[name] = []header.Header{h}
		a.order = append(a.order, name)
	}
	a.mu.Unlock()
}

func (a *headers) String() string {
	buffer := bytes.Buffer{}
	a.mu.Lock()

	for i, name := range a.order {
		hds := a.headers[name]

		for idx, h := range hds {
			buffer.WriteString(h.String())
			if i < len(a.order) || idx < len(hds) {
				buffer.WriteString("\r\n")
			}
		}
	}

	a.mu.Unlock()
	return buffer.String()
}

func (a *headers) PrependHeader(h header.Header) {
	name := strings.ToLower(h.Name())
	a.mu.Lock()
	if hd, ok := a.headers[name]; ok {
		a.headers[name] = append([]header.Header{h}, hd...)
	} else {
		a.headers[name] = []header.Header{h}
		newOrder := make([]string, 1, len(a.order)+1)
		newOrder[0] = name
		a.order = append(newOrder, a.order...)
	}
	a.mu.Unlock()
}

func (a *headers) PrePendHeaderAfter(h header.Header, afterName string) {
	headerName := strings.ToLower(h.Name())
	afterName = strings.ToLower(afterName)
	a.mu.Lock()
	if _, ok := a.headers[afterName]; ok {
		afterIdx := -1
		headerIdx := -1
		for i, name := range a.order {
			if name == afterName {
				afterIdx = i
			}
			if name == headerName {
				headerIdx = i
			}
		}

		if headerIdx == -1 {
			a.headers[headerName] = []header.Header{h}
			newOrder := make([]string, 0)
			newOrder = append(newOrder, a.order[:afterIdx+1]...)
			newOrder = append(newOrder, headerName)
			newOrder = append(newOrder, a.order[afterIdx+1:]...)
			a.order = newOrder
		} else {
			a.headers[headerName] = append([]header.Header{h}, a.headers[headerName]...)
			newOrder := make([]string, 0)
			if afterIdx < headerIdx {
				newOrder = append(newOrder, a.order[:afterIdx+1]...)
				newOrder = append(newOrder, headerName)
				newOrder = append(newOrder, a.order[afterIdx+1:headerIdx]...)
				newOrder = append(newOrder, a.order[headerIdx+1:]...)
			} else {
				newOrder = append(newOrder, a.order[:headerIdx]...)
				newOrder = append(newOrder, a.order[headerIdx+1:afterIdx+1]...)
				newOrder = append(newOrder, headerName)
				newOrder = append(newOrder, a.order[afterIdx+1:]...)
			}
			a.order = newOrder
		}
		a.mu.Unlock()
	} else {
		a.mu.Unlock()
		a.PrependHeader(h)
	}
}

func (a *headers) ReplaceHeader(name string, hds []header.Header) {
	name = strings.ToLower(name)
	a.mu.Lock()
	if _, ok := a.headers[name]; ok {
		a.headers[name] = hds
	}
	a.mu.Unlock()
}

func (a *headers) Headers() []header.Header {
	hds := make([]header.Header, 0)

	a.mu.Lock()
	defer a.mu.Unlock()
	for _, v := range a.headers {
		for _, h := range v {
			hds = append(hds, h)
		}
	}

	return hds
}

func (a *headers) GetHeaders(name string) []header.Header {
	name = strings.ToLower(name)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.headers != nil {
		if hds, ok := a.headers[name]; ok {
			return hds
		}
	}
	return []header.Header{}
}

func (a *headers) Remove(name string)  {
	name = strings.ToLower(name)
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.headers != nil {
		delete(a.headers, name)
	}

	for i, entry := range a.order {
		if entry  == name {
			a.order = append(a.order[:i], a.order[i+1:]...)
		}
	}
}

// CloneHeaders returns all cloned headers in slice.
func (a *headers) CloneHeaders() []header.Header {
	return cloneHeaders(a)
}

func cloneHeaders(msg interface{ Headers() []header.Header }) []header.Header {
	hdrs := make([]header.Header, 0)
	for _, h := range msg.Headers() {
		hdrs = append(hdrs, h.Clone())
	}
	return hdrs
}

func (a *headers) ContentLength() (*header.ContentLength, bool) {
	hdrs := a.GetHeaders("Content-Length")
	if len(hdrs) == 0 {
		return nil, false
	}
	contentLength, ok := hdrs[0].(*header.ContentLength)
	if !ok {
		return nil, false
	}
	return contentLength, true
}

