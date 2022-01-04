package header

import (
	"fmt"
	"strings"
)

type ProxyRequire struct {
	Options []string
}

func (r *ProxyRequire) String() string {
	return fmt.Sprintf("%s: %s", r.Name(), r.Value())
}

func (r *ProxyRequire) Name() string {
	return "Proxy-Require"
}

func (r *ProxyRequire) Value() string {
	return strings.Join(r.Options, ", ")
}

func (r *ProxyRequire) Equals(other interface{}) bool {
	if h, ok := other.(*ProxyRequire); ok {
		if r == h {
			return true
		}
		if h == nil && r != nil || h != nil && r == nil {
			return false
		}

		if len(r.Options) != len(h.Options) {
			return false
		}

		for i, opt := range r.Options {
			if opt != h.Options[i] {
				return false
			}
		}

		return true
	}
	return false
}

func (r *ProxyRequire) Clone() Header {
	if r == nil {
		var newProxyRequire *ProxyRequire
		return newProxyRequire
	}

	dup := make([]string, len(r.Options))
	copy(dup, r.Options)
	return &ProxyRequire{dup}
}
