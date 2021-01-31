package message

import "time"

type ID [16]byte

// Message
type Message struct {
	Meta    *Metadata   // meta data
	Message interface{} // message body
}

//Metadata The message metadata
type Metadata struct {
	ID        ID        // message ID
	Time      time.Time // message time
	ServiceID int    	// message product service ID
	Host      string    // message product server
}

func (m *Message) Parse(buf []byte) (err error) {
	m.Meta = &Metadata{}
	copy(m.Meta.ID[:], buf[:15])
	return
}
