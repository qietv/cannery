package pkg

import (
	"os"
)

type Project struct {
	ID     string
	Name   string
	Path   string
	Offset int
	Status Status
	Latest *os.File
}
