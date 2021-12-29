package sip

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

type Params interface {
	Get(key string) (Stringer, bool)
	Add(key string, val Stringer) Params
	Remove(key string) Params
	Clone() Params
	Equals(params interface{}) bool
	ToString(sep uint8) string
	String() string
	Length() int
	Items() map[string]Stringer
	Keys() []string
	Has(key string) bool
}

type headerParams struct {
	mu          sync.Mutex
	params      map[string]Stringer
	paramsOrder []string
}

func NewParams() Params {
	return &headerParams{
		params:      make(map[string]Stringer),
		paramsOrder: []string{},
	}
}

func (h *headerParams) Get(key string) (Stringer, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	p, ok := h.params[key]
	return p, ok
}

func (h *headerParams) Add(key string, val Stringer) Params {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.params[key]; !ok {
		h.paramsOrder = append(h.paramsOrder, key)
	}

	h.params[key] = val
	return h
}

func (h *headerParams) Remove(key string) Params {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.params[key]; ok {
		delete(h.params, key)

		for k, v := range h.paramsOrder {
			if v == key {
				h.paramsOrder = append(h.paramsOrder[:k], h.paramsOrder[k+1:]...)
				break
			}
		}
	}
	return h
}

func (h *headerParams) Clone() Params {
	if h == nil {
		var dup *headerParams
		return dup
	}

	dup := NewParams()
	for _, key := range h.Keys() {
		if val, ok := h.Get(key); ok {
			dup.Add(key, val)
		}
	}

	return dup
}

func (h *headerParams) Equals(other interface{}) bool {
	q, ok := other.(*headerParams)
	if !ok {
		return false
	}

	if h == q {
		return true
	}
	if h == nil && q != nil || h != nil && q == nil {
		return false
	}

	if h.Length() == 0 && q.Length() == 0 {
		return true
	}

	if h.Length() != q.Length() {
		return false
	}

	for key, pVal := range h.Items() {
		qVal, ok := q.Get(key)
		if !ok {
			return false
		}
		if pVal != qVal {
			return false
		}
	}

	return true
}

func (h *headerParams) ToString(sep uint8) string {
	if h == nil {
		return ""
	}

	var buffer bytes.Buffer
	first := true

	for _, key := range h.Keys() {
		val, ok := h.Get(key)
		if !ok {
			continue
		}

		if !first {
			buffer.WriteString(fmt.Sprintf("%c", sep))
		}
		first = false

		buffer.WriteString(fmt.Sprintf("%s", key))

		if val, ok := val.(*String); ok {
			if strings.ContainsAny(val.String(), " \t") {
				buffer.WriteString(fmt.Sprintf("=\"%s\"", val.String()))
			} else {
				buffer.WriteString(fmt.Sprintf("=%s", val.String()))
			}
		}
	}

	return buffer.String()
}

func (h *headerParams) String() string {
	if h == nil {
		return ""
	}

	return h.ToString('&')
}

func (h *headerParams) Length() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	return len(h.params)
}

func (h *headerParams) Items() map[string]Stringer {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.params
}

func (h *headerParams) Keys() []string {
	h.mu.Lock()
	defer h.mu.Unlock()

	return h.paramsOrder
}

func (h *headerParams) Has(key string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	_, ok := h.params[key]
	return ok
}
