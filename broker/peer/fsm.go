package peer

import (
	"bufio"
	"encoding/json"
	"github.com/hanskorg/logkit"
	"github.com/hashicorp/raft"
	"io"
	"sync"
)

type Event struct {
	EventName string `yaml:"event"`
	LogName   string `yaml:"key"`
	LogBody   string `yaml:"value"`
}

type fsmSnapshot struct {
	store sync.Map
}

type fsm struct {
	mu    sync.Mutex
	store sync.Map
}

func (fsm *fsm) Apply(l *raft.Log) interface{} {
	var (
		err error
		evt *Event
	)
	evt = &Event{}
    err = json.Unmarshal(l.Data, evt)
    if err != nil {
        logkit.Warnf("fsm: apply unmarshal fail, %v", err.Error())
        return nil
    }
    switch evt.EventName {
    case "PUT","POST":
        fsm.store.Store(evt.LogName, evt.LogBody)
    case "DELETE":
        fsm.store.Delete(evt.LogName)
    }
    return nil
}

func (fsm *fsm) Snapshot() (fs raft.FSMSnapshot, err error) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	fs = &fsmSnapshot{
		store: sync.Map{},
	}
	fsm.store.Range(func(key, value interface{}) bool {
		fs.(*fsmSnapshot).store.Store(key, value)
		return true
	})
	return

}

func (fsm *fsm) Restore(reader io.ReadCloser) (err error) {
	var (
		rd *bufio.Reader
	)
	fsm.store = sync.Map{}
	rd = bufio.NewReader(reader)
	for line, _, e := rd.ReadLine(); e != io.EOF; {
		log := map[string]string{}
		if e != nil {
			logkit.Warnf("fsm: restore read line fail, %v", e.Error())
			continue
		}
		e = json.Unmarshal(line, &log)
		if e != nil {
			logkit.Warnf("fsm: restore unmarshal line fail, %v", e.Error())
			continue
		}
		fsm.store.Store(log["k"], log["v"])
	}
	return
}

func (fsmS *fsmSnapshot) Persist(sink raft.SnapshotSink) (err error) {

	err = func() (err error) {
		fsmS.store.Range(func(key, value interface{}) bool {
			var (
				buff []byte
			)
			buff, err = json.Marshal(fsmS.store)
			if err != nil {
				return false
			}
			if _, err = sink.Write(buff); err != nil {
				return false
			}
			sink.Write([]byte("\n"))
			return true
		})
		return sink.Close()
	}()
	if err != nil {
		sink.Cancel()
	}
	return
}

func (fsmS *fsmSnapshot) Release() {
	logkit.Infof("fsm: snap release")
}
