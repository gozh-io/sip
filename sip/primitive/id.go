package primitive

import "github.com/google/uuid"

type MessageID string

func NextMessageID() MessageID {
	return MessageID(uuid.Must(uuid.NewUUID()).String())
}
