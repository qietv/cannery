package agent

type STATUS int

const (
	_ STATUS = iota
	PREPARE
	RUNNING
	CRASHED
	PAUSED
	COMPLETED
)
type Status struct {
	Signal chan int
	Projects map[string]*Project
	_status STATUS
}

func (s *Status) Get() STATUS  {
	return s._status
}
