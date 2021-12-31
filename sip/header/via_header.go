package header

import (
	"bytes"
	"fmt"

	"github.com/gozh-io/sip/sip/primitive"
	"github.com/gozh-io/sip/util"
)

type ViaHop struct {
	ProtocolName    string
	ProtocolVersion string
	Transport       string
	Host   string
	Port   *primitive.Port
	Params Params
}

type ViaHeader []*ViaHop

func NewViaHop(name, version, transport string, host string, port *primitive.Port, params Params) *ViaHop {
	return &ViaHop{
		ProtocolName:    name,
		ProtocolVersion: version,
		Transport:       transport,
		Host:            host,
		Port:            port,
		Params:          params,
	}
}

func (v ViaHeader) Equals(other interface{}) bool {
	if h, ok := other.(ViaHeader); ok {
		if len(v) != len(h) {
			return false
		}

		for i, hop := range v {
			if !hop.Equals(h[i]) {
				return false
			}
		}

		return true
	}

	return false
}

func (v ViaHeader) Clone() ViaHeader {
	if v == nil {
		var newHeader ViaHeader
		return newHeader
	}

	dup := make([]*ViaHop, 0)

	for _, hop := range v {
		dup = append(dup, hop.Clone())
	}

	return dup
}

func (v ViaHeader) String() string {
	return fmt.Sprintf("%s %s", v.Name(), v.Value())
}

func (v ViaHeader) Name() string {
	return "Via"
}

func (v ViaHeader) Value() string {
	var buffer bytes.Buffer
	for idx, hop := range v {
		buffer.WriteString(hop.String())
		if idx != len(v)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

func (v *ViaHop) SentBy() string {
	var buf bytes.Buffer
	buf.WriteString(v.Host)
	if v.Port != nil {
		buf.WriteString(fmt.Sprintf(":%d", *v.Port))
	}

	return buf.String()
}

func (v ViaHop) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%s/%s/%s %s", v.ProtocolName, v.ProtocolVersion, v.Transport, v.Host))

	if v.Port != nil {
		buffer.WriteString(fmt.Sprintf(":%d", *v.Port))
	}

	if v.Params != nil && v.Params.Length() > 0 {
		buffer.WriteString(";")
		buffer.WriteString(v.Params.ToString(';'))
	}

	return buffer.String()
}

func (v *ViaHop) Clone() *ViaHop {
	var newHop *ViaHop

	newHop = &ViaHop{
		ProtocolName:    v.ProtocolName,
		ProtocolVersion: v.ProtocolVersion,
		Transport:       v.Transport,
		Host:            v.Host,
	}

	if v.Port != nil {
		newHop.Port = v.Port
	}
	if v.Params != nil {
		newHop.Params = v.Params
	}

	return newHop
}

func (v *ViaHop) Equals(other interface{}) bool {
	if h, ok := other.(*ViaHop); ok {
		if v == h {
			return true
		}
		if v == nil && h != nil || v != nil && h == nil {
			return false
		}

		res := v.ProtocolName == h.ProtocolName &&
			v.ProtocolVersion == h.ProtocolVersion &&
			v.Transport == h.Transport &&
			v.Host == h.Host &&
			util.Uint16PtrEq((*uint16)(v.Port), (*uint16)(h.Port))

		if v.Params != h.Params {
			if v.Params == nil {
				res = res && h.Params == nil
			} else {
				res = res && v.Params.Equals(h.Params)
			}
		}

		return res
	}

	return false
}
