package primitive

type RequestMethod int

const (
	INVITE RequestMethod = iota
	ACK
	CANCEL
	BYE
	REGISTER
	OPTIONS
	SUBSCRIBE
	NOTIFY
	REFER
	INFO
	MESSAGE
	PRACK
	UPDATE
	PUBLISH
)

func (rm RequestMethod) String() string {
	var m string
	switch rm {
	case INVITE:
		m = "INVITE"
	case BYE:
		m = "BYE"
	case ACK:
		m = "ACK"
	case CANCEL:
		m = "CANCEL"
	case REGISTER:
		m = "REGISTER"
	case OPTIONS:
		m = "OPTIONS"
	case SUBSCRIBE:
		m = "SUBSCRIBE"
	case PUBLISH:
		m = "PUBLISH"
	case NOTIFY:
		m = "NOTIFY"
	case MESSAGE:
		m = "MESSAGE"
	case REFER:
		m = "REFER"
	case INFO:
		m = "INFO"
	case PRACK:
		m = "PRACK"
	case UPDATE:
		m = "UPDATE"
	}
	return m
}
