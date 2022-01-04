package header

import (
	"fmt"
)

type Subject struct {
	Sender                 string
	SenderEndPointSerial   uint16
	Receiver               string
	ReceiverEndPointSerial uint16
}

func (s *Subject) Name() string {
	return "Subject"
}

func (s *Subject) Value() string {
	return fmt.Sprintf("%s:%d,%s:%d", s.Sender, s.SenderEndPointSerial, s.Receiver, s.ReceiverEndPointSerial)
}

func (s *Subject) Clone() Header {
	return s
}

func (s *Subject) String() string {
	return fmt.Sprintf("%s: %s", s.Name(), s.Value())
}

func (s *Subject) Equals(other interface{}) bool {
	if h, ok := other.(*Subject); ok {
		if h == s {
			return true
		}
		return false
	}
	return false
}
