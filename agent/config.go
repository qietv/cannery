package agent

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	ServerID      int    `yaml:"serverId" json:"serverId"`
	ServerName    string `yaml:"serverName" json:"serverName"`
	Listener      string `yaml:"listener" json:"listener"`
	RemoteServer  string `yaml:"remoter" json:"remoter"`
	CommitTimeout time.Time
	Recoverable   bool
	Database      string //sqlite、kv

}

func DefaultConfig() *Config {
	return &Config{
		ServerID:      0,
		Listener:      "",
		RemoteServer:  "",
		CommitTimeout: time.Time{},
		Recoverable:   false,
	}
}
func ValidateConfig(c *Config) (err error) {
	if c.Listener == "" || (!strings.HasPrefix(c.Listener, "grpc://") && strings.HasPrefix(c.Listener, "http://")) {
		return fmt.Errorf("listener must be prefix: grpc:// or http://")
	}
	if c.ServerID == 0 {
		return fmt.Errorf("serverId must be a num that bigger 0")
	}
	return
}
