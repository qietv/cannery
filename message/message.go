package message

import "time"

type Message struct {
	Meta    *Metadata              // meta data
	Message map[string]interface{} // message body
}

type Metadata struct {
	ID        []byte    // message ID
	Host      string    // message product server
	ServiceID string    // message product service ID
	MTime     time.Time // message time
	TrackId   string
}
