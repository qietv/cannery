package peer

import (
	"encoding/json"
	"github.com/hanskorg/logkit"
	"github.com/hashicorp/raft"
	"io"
	"sync"
)

type fsmSnapshot struct {
	store map[string]string
}

type fsm struct {
	mu    sync.Mutex
	store map[string]string
}

func (fsm *fsm) Apply(l *raft.Log) interface{} {
	var (
		err error
	)
	l.Type
}

func (fsm *fsm) Snapshot() (fs raft.FSMSnapshot, err error) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	fs = &fsmSnapshot{
		store: map[string]string{},
	}
	for key, value := range fsm.store {
        fs.(*fsmSnapshot).store[key] = value
    }
	return

}

func (fsm *fsm) Restore(reader io.ReadCloser) (err error) {
	fsm.store = map[string]string{}
    if err = json.NewDecoder(reader).Decode(fsm.store); err != nil {
        return err
    }
    return
}

func (fsmS *fsmSnapshot) Persist(sink raft.SnapshotSink) (err error) {

    err = func() (err error){
        var (
            buff []byte
        )
        buff, err = json.Marshal(fsmS.store)
        if err != nil {
            return err
        }
        if _, err = sink.Write(buff); err != nil {
            return err
        }
        return sink.Close()
    }()
    if err != nil {
        sink.Cancel()
    }
    return
}

func (fsmS *fsmSnapshot) Release() {
	logkit.Infof("snap finish, %d", len(fsmS.store))
}
