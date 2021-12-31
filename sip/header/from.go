package header

import (
	"bytes"
	"fmt"

	"github.com/gozh-io/sip/sip/primitive"
)

type FromHeader struct {
	DisplayName primitive.Stringer
	Address     Uri
	Params      Params
}

func (f *FromHeader) String() string {
	return fmt.Sprintf("%s: %s", f.Name(), f.Value())
}

func (f *FromHeader) Name() string {
	return "From"
}

func (f *FromHeader) Value() string {
	var buffer bytes.Buffer

	if name, ok := f.DisplayName.(*primitive.String); ok && name.String() != "" {
		buffer.WriteString(fmt.Sprintf("\"%s\"", name))
	}

	buffer.WriteString(fmt.Sprintf("<%s>", f.Address))
	if f.Params != nil && f.Params.Length() > 0 {
		buffer.WriteString(";")
		buffer.WriteString(f.Params.String())
	}

	return buffer.String()
}

func (f *FromHeader) Equals(other interface{}) bool {
	if h, ok := other.(*FromHeader); ok {
		if f == h {
			return true
		}
		if f == nil && h != nil || f != nil && h == nil {
			return false
		}

		res := true

		if f.DisplayName != h.DisplayName {
			if f.DisplayName == nil {
				res = res && h.DisplayName == nil
			} else {
				res = res && f.DisplayName.Equals(h.DisplayName)
			}
		}

		if f.Address != h.Address {
			if f.Address == nil {
				res = res && h.Address == nil
			} else {
				res = res && f.Address.Equals(h.Address)
			}
		}

		if f.Params != h.Params {
			if f.Params == nil {
				res = res && h.Params == nil
			} else {
				res = res && f.Params.Equals(h.Params)
			}
		}

		return res
	}

	return false
}

func (f *FromHeader) Clone() Header {
	var newTo *ToHeader
	if f == nil {
		return newTo
	}

	newTo = &ToHeader{
		DisplayName: f.DisplayName,
	}
	if f.Address != nil {
		newTo.Address = f.Address.Clone()
	}
	if f.Params != nil {
		newTo.Params = f.Params.Clone()
	}
	return newTo
}
