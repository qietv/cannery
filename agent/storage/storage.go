package storage

import (
	"context"
	"errors"
)

var (
	UnknownProject = errors.New("project unknown")
)

type Offset struct {
	ID      int    `json:"_"`
	Project string `json:"project"`
	Offset  int    `json:"offset"`
}

type Project struct {
	ID       int    `json:"_"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Hostname string `json:"host"`
	Status   string `json:"status"`
}

type Storage interface {
	//AddProject add new project for event collect
	AddProject(ctx context.Context, host, project, path string) (err error)
	//RmProject remove project from event bus
	RemoveProject(ctx context.Context, project string) (err error)
	//Read latest read line
	Read(ctx context.Context, project string) (offset int, err error)
	//UpdateRead commit read offset
	UpdateRead(ctx context.Context, project string, offset int, replace bool) (err error)
}
