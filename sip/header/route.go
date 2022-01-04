package header

import (
	"fmt"
	"strings"
)

type Route struct {
	Address []Uri
}


func (r *Route) String() string {
	return fmt.Sprintf("%s: %s", r.Name(), r.Value())
}

func (r *Route) Name() string {
	return "Route"
}

func (r *Route) Value() string {
	var addrs []string
	for _, uri := range r.Address {
		addrs = append(addrs, "<"+ uri.String() + ">")
	}

	return  strings.Join(addrs, ", ")
}

func (r *Route) Clone() Header {
	var newRoute *Route
	if r == nil {
		return newRoute
	}

	newRoute = &Route{
		Address: make([]Uri, len(r.Address)),
	}

	for i, uri := range r.Address {
		newRoute.Address[i] = uri.Clone()
	}

	return newRoute
}


func (r *Route) Equals(other interface{}) bool {
	if h, ok := other.(*Route); ok {
		if r == h {
			return true
		}
		if r == nil && h != nil || r != nil && h == nil {
			return false
		}

		for i, uri := range r.Address {
			if !uri.Equals(h.Address[i]) {
				return false
			}
		}

		return true
	}

	return false
}

