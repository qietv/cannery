package config

import (
	. "github.com/qietv/qgrpc/pkg"
	"time"
)

type Config struct {
	ServerID          string   `yaml:"serverId"`
	DataDir           string   `yaml:"dataDir"`
	CommitTimeout     Duration `yaml:"commitTimeout"`
	SnapshotDir       string   `yaml:"snapshotPath"`
	SnapshotRetain    int      `yaml:"snapshotRetain"`
	SnapshotThreshold uint64   `yaml:"snapshotThreshold"`
	SnapshotInterval  Duration `yaml:"snapshotInterval"`
	MaxAppendEntries  int      `yaml:"MaxAppendEntries"`
	TrailingLogs      uint64   `yaml:"trailingLogs"`
	Peer              *Peer    `yaml:"peer"`
}
type Peer struct {
	Listener  string   `yaml:"listener"`
	Advertise string   `yaml:"advertise"`
	Timeout   *Timeout `yaml:"timeout"`
	LogLevel  string   `yaml:"logLevel"`
}

type Timeout struct {
	Commit      Duration `yaml:"commit"`
	Transport   Duration `yaml:"transport"`
	Heartbeat   Duration `yaml:"heartbeat"`
	Election    Duration `yaml:"election"`
	LeaderLease Duration `yaml:"leaderLease"`
}

func DefaultConfig() *Config {
	return &Config{
		DataDir:           "/var/run/cannery/data/",
		CommitTimeout:     Duration(50 * time.Millisecond),
		SnapshotDir:       "/var/run/cannery/snapshot/",
		SnapshotRetain:    5,
		MaxAppendEntries:  64,
		TrailingLogs:      10240,
		SnapshotThreshold: 8192,
		SnapshotInterval:  Duration(120 * time.Second),
		Peer: &Peer{
			Listener: "0.0.0.0:8808",
			Timeout: &Timeout{
				Transport:   Duration(10 * time.Second),
				Heartbeat:   Duration(1000 * time.Millisecond),
				Election:    Duration(1000 * time.Millisecond),
				LeaderLease: Duration(500 * time.Millisecond),
			},
			LogLevel: "INFO",
		},
	}
}
