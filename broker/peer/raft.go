package peer

import (
	"cannery/broker/config"
	"errors"
	"github.com/hanskorg/logkit"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"io"
	"log"
	"net"
	"path/filepath"
	"time"
)

var (
	NotLeaderErr = errors.New("not leader node")
	NotFoundErr  = errors.New("that key not exists")
)

type Options struct {
	HeartbeatTimeout  time.Duration
	CommitTimeout     time.Duration
	SnapshotInterval  time.Duration
	ElectionTimeout   time.Duration
	SnapshotThreshold uint64
}
type Raftd struct {
	store   *fsm
	raft    *raft.Raft
	options *Options
}

func NewRaftd(c *config.Config, logWriter io.Writer) *Raftd {
	var (
		rc                *raft.Config
		peerAdvertiseAddr net.Addr
		peerTransport     *raft.NetworkTransport
		snapshotStore     *raft.FileSnapshotStore
		logStore          raft.LogStore
		stableStore       raft.StableStore
		s                 *fsm
		err               error
	)
	rc = &raft.Config{
		ProtocolVersion:    raft.ProtocolVersionMax,
		HeartbeatTimeout:   time.Duration(c.Peer.Timeout.Heartbeat),
		ElectionTimeout:    time.Duration(c.Peer.Timeout.Election),
		CommitTimeout:      time.Duration(c.Peer.Timeout.Commit),
		MaxAppendEntries:   c.MaxAppendEntries,
		ShutdownOnRemove:   true,
		TrailingLogs:       c.TrailingLogs,
		SnapshotInterval:   time.Duration(c.SnapshotInterval),
		SnapshotThreshold:  c.SnapshotThreshold,
		LeaderLeaseTimeout: time.Duration(c.Peer.Timeout.LeaderLease),
		LocalID:            raft.ServerID(c.ServerID),
		LogOutput:          logWriter,
		LogLevel:           "",
		Logger:             nil,
	}
	peerAdvertiseAddr, err = net.ResolveTCPAddr("tcp", c.Peer.Advertise)
	if err != nil {
		log.Fatalf("raft: peer can not listen, %s", err.Error())
	}
	peerTransport, err = raft.NewTCPTransport(c.Peer.Listener, peerAdvertiseAddr, 10, time.Duration(c.Peer.Timeout.Transport), logWriter)
	if err != nil {
		log.Fatalf("raft: peer transport init fail, %s", err.Error())
	}
	snapshotStore, err = raft.NewFileSnapshotStore(c.SnapshotDir, c.SnapshotRetain, logWriter)
	if err != nil {
		log.Fatalf("raft: snapshot init fail, %s", err.Error())
	}
	logStore = raft.NewInmemStore()
	stableStore, err = raftboltdb.NewBoltStore(filepath.Join(c.DataDir, "stable.db"))
	if err != nil {
		log.Fatalf("raft: stable init fail, %s", err.Error())
	}
	s = &fsm{}
	r, err := raft.NewRaft(rc, s, logStore, stableStore, snapshotStore, peerTransport)
	if err != nil {
		log.Fatalf("raft: raft start fail, %v", err.Error())
	}

	configuration := raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      raft.ServerID(c.ServerID),
				Address: peerTransport.LocalAddr(),
			},
		},
	}
	r.BootstrapCluster(configuration)

	return &Raftd{
		raft: r,
		options: &Options{
			HeartbeatTimeout:  time.Duration(c.Peer.Timeout.Heartbeat),
			ElectionTimeout:   time.Duration(c.Peer.Timeout.Election),
			CommitTimeout:     time.Duration(c.Peer.Timeout.Commit),
			SnapshotInterval:  time.Duration(c.SnapshotInterval),
			SnapshotThreshold: c.SnapshotThreshold,
		},
		store: s,
	}
}

func (p *Raftd) Join(nodeID, addr string) (err error) {
	var (
		cf          raft.ConfigurationFuture
		indexFuture raft.IndexFuture
	)
	cf = p.raft.GetConfiguration()
	if cf.Error() != nil {
		return
	}
	for _, srv := range cf.Configuration().Servers {
		if srv.ID == raft.ServerID(nodeID) || srv.Address == raft.ServerAddress(addr) {
			if srv.ID == raft.ServerID(nodeID) && srv.Address == raft.ServerAddress(addr) {
				logkit.Infof("node %s at %s already member of cluster, ignoring join request", nodeID, addr)
				return
			}
		}
		p.raft.RemoveServer(srv.ID, 0, 0)
		if err := cf.Error(); err != nil {
			logkit.Errorf("error removing existing node %s at %s: %s", nodeID, addr, err)
			return
		}
	}
	indexFuture = p.raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if err = indexFuture.Error(); err != nil {
		logkit.Errorf("node %s add voter fail at %s: %s", nodeID, addr, err)
		return
	}
	logkit.Infof("node %s at %s joined successfully", nodeID, addr)
	return
}

func (p *Raftd) Delete(key string) error {
	if p.raft.State() != raft.Leader {
		return NotLeaderErr
	}
	p.raft.Apply((&Event{
		EventName: "DELETE",
		LogName:   key,
		LogBody:   "",
	}).toBytes(), p.options.CommitTimeout)
	return nil
}
func (p *Raftd) Post(key string, val interface{}) (err error) {
	if p.raft.State() != raft.Leader {
		return NotLeaderErr
	}
	p.raft.Apply((&Event{
		EventName: "POST",
		LogName:   key,
		LogBody:   val,
	}).toBytes(), p.options.CommitTimeout)

	return
}

func (p *Raftd) Get(key string) (value interface{}, err error) {
	var has bool
	value, has = p.store.store.Load(key)
	if !has {
		err = NotFoundErr
	}
	return
}

func (p *Raftd) isLeader() bool {
	return p.raft.State() == raft.Leader
}
