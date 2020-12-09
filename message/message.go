package message

import "time"

type Message struct {
	Meta    *Metadata              // meta data
	Message map[string]interface{} // message body
}

type Metadata struct {
	ID        [16]byte    // message ID
	Host      string    // message product server
	ServiceID string    // message product service ID
	Time      time.Time // message time
	TrackId   string
}
