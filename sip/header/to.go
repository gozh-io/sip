package header

import (
	"bytes"
	"fmt"

	"github.com/gozh-io/sip/sip/primitive"
)

type ToHeader struct {
	DisplayName primitive.Stringer
	Address     Uri
	Params      Params
}

func (t *ToHeader) String() string {
	return fmt.Sprintf("%s: %s", t.Name(), t.Value())
}

func (t *ToHeader) Name() string {
	return "To"
}

func (t *ToHeader) Value() string {
	var buffer bytes.Buffer

	if name, ok := t.DisplayName.(*primitive.String); ok && name.String() != "" {
		buffer.WriteString(fmt.Sprintf("\"%s\"", name))
	}

	buffer.WriteString(fmt.Sprintf("<%s>", t.Address))

	if t.Params != nil && t.Params.Length() > 0 {
		buffer.WriteString(";")
		buffer.WriteString(t.Params.String())
	}

	return buffer.String()
}

func (t *ToHeader) Equals(other interface{}) bool {
	if h, ok := other.(*ToHeader); ok {
		if t == h {
			return true
		}
		if t == nil && h != nil || t != nil && h == nil {
			return false
		}

		res := true

		if t.DisplayName != h.DisplayName {
			if t.DisplayName == nil {
				res = res && h.DisplayName == nil
			} else {
				res = res && t.DisplayName.Equals(h.DisplayName)
			}
		}

		if t.Address != h.Address {
			if t.Address == nil {
				res = res && h.Address == nil
			} else {
				res = res && t.Address.Equals(h.Address)
			}
		}

		if t.Params != h.Params {
			if t.Params == nil {
				res = res && h.Params == nil
			} else {
				res = res && t.Params.Equals(h.Params)
			}
		}

		return res
	}

	return false
}

func (t *ToHeader) Clone() Header {
	var newTo *ToHeader
	if t == nil {
		return newTo
	}

	newTo = &ToHeader{
		DisplayName: t.DisplayName,
	}
	if t.Address != nil {
		newTo.Address = t.Address.Clone()
	}
	if t.Params != nil {
		newTo.Params = t.Params.Clone()
	}
	return newTo
}
