package header

import (
	"fmt"
	"strings"
)

type RecordRouteHeader struct {
	Addresses []Uri
}

func (route *RecordRouteHeader) Name() string { return "Record-Route" }

func (route *RecordRouteHeader) Value() string {
	var addrs []string
	for _, uri := range route.Addresses {
		addrs = append(addrs, "<"+uri.String()+">")
	}
	return strings.Join(addrs, ", ")
}

func (route *RecordRouteHeader) String() string {
	return fmt.Sprintf("%s: %s", route.Name(), route.Value())
}

func (route *RecordRouteHeader) Clone() Header {
	var newRoute *RecordRouteHeader
	if route == nil {
		return newRoute
	}

	newRoute = &RecordRouteHeader{
		Addresses: make([]Uri, len(route.Addresses)),
	}

	for i, uri := range route.Addresses {
		newRoute.Addresses[i] = uri.Clone()
	}

	return newRoute
}

func (route *RecordRouteHeader) Equals(other interface{}) bool {
	if h, ok := other.(*RecordRouteHeader); ok {
		if route == h {
			return true
		}
		if route == nil && h != nil || route != nil && h == nil {
			return false
		}

		for i, uri := range route.Addresses {
			if !uri.Equals(h.Addresses[i]) {
				return false
			}
		}

		return true
	}

	return false
}
