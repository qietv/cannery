package peer

import "encoding/json"

type Peer interface {
	Delete(key string) error

	Post(key string, val interface{}) error

	Get(key string) (val interface{}, err error)
}

type Event struct {
	EventName string      `yaml:"event"`
	LogName   string      `yaml:"key"`
	LogBody   interface{} `yaml:"value"`
}

func (e *Event) toBytes() (out []byte) {
	out, _ = json.Marshal(e)
	return
}

func (e *Event) Unmarshal(buf []byte) (err error) {
	err = json.Unmarshal(buf, e)
	return
}
