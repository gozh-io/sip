package sip

type TxKey string

func (tk TxKey) String() string {
	return string(tk)
}

type TX interface {
	Origin() Request
	Key() TxKey
	String() string
	Errors() <-chan error
	Done() <-chan bool
}

type ClientTx interface {
	TX
	Responses() <-chan Response
	Cancel() error
}

type ServerSTx interface {
	TX
	Respond(res Response) error
	Acks() <-chan Request
	Cancels() <-chan Request
}
