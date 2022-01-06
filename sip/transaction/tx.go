package transaction

import (
	"log"
	"sync"

	"github.com/gozh-io/sip/sip"
)

type TxKey = sip.TxKey

type Tx interface {
	Init() error
	Key() TxKey
	Origin() sip.Request
	Receive(msg sip.Message) error
	String() string
	Transport() sip.Transport
	Terminate()
	Errors() <-chan error
	Done() <-chan bool
}

type commonTx struct {
	key          TxKey
	fsm          *sip.FSM
	fsmMu        sync.RWMutex
	origin       sip.Request
	tp           sip.Transport
	lastResponse sip.Response

	errs    chan error
	lastErr error
	done    chan bool
	log     log.Logger
}
