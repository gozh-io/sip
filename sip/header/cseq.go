package header

import (
	"fmt"

	"github.com/gozh-io/sip/sip/primitive"
)

type CSeq struct {
	SeqNo      uint32
	MethodName primitive.RequestMethod
}

func (a *CSeq) String() string {
	return fmt.Sprintf("%s: %s", a.Name(), a.Value())
}

func (a *CSeq) Name() string {
	return "CSeq"
}

func (a *CSeq) Value() string {
	return fmt.Sprintf("%d %s", a.SeqNo, a.MethodName.String())
}

func (a *CSeq) Equals(other interface{}) bool {
	if h, ok := other.(*CSeq); ok {
		if a == h {
			return true
		}
		if a == nil && h != nil || a != nil && h == nil {
			return false
		}

		return a.SeqNo == h.SeqNo &&
			a.MethodName == h.MethodName
	}

	return false
}

func (a *CSeq) Clone() Header {
	if a == nil {
		var newCSeq *CSeq
		return newCSeq
	}

	return &CSeq{
		SeqNo:      a.SeqNo,
		MethodName: a.MethodName,
	}
}
