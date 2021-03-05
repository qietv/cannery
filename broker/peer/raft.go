package peer

import (
    "cannery/broker/config"
    "github.com/hashicorp/raft"
    "github.com/hashicorp/raft-boltdb"
    "io"
    "log"
    "net"
    "path/filepath"
    "time"
)

func NewRaft(c *config.Config, logWriter io.Writer) {
	var (
		rc                *raft.Config
		peerAdvertiseAddr net.Addr
		peerTransport     *raft.NetworkTransport
		snapshotStore     *raft.FileSnapshotStore
		logStore          raft.LogStore
		stableStore       raft.StableStore
		err               error
	)
	rc = &raft.Config{
		ProtocolVersion:    raft.ProtocolVersionMax,
		HeartbeatTimeout:   time.Duration(c.Peer.Timeout.Heartbeat),
		ElectionTimeout:    time.Duration(c.Peer.Timeout.Heartbeat),
		CommitTimeout:      time.Duration(c.Peer.Timeout.Heartbeat),
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

    r, err := raft.NewRaft(rc, &fsm{}, logStore, stableStore, snapshotStore, peerTransport)
    if err != nil {
        log.Fatalf("raft: raft start fail, %v", err.Error())
    }

    configuration := raft.Configuration{
        Servers: []raft.Server{
            {
                ID:       raft.ServerID(c.ServerID),
                Address:  peerTransport.LocalAddr(),
            },
        },
    }

    r.BootstrapCluster(configuration)
    r.Leader()
}
