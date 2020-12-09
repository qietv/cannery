package agent

import (
    "github.com/hpcloud/tail"
    "strings"
)

var (
	agent *logAgent
)

type logAgent struct {
	tailMan *tail.Tail
	offset  int64
}

//NewTail new log tail instance
func NewTail(file string) (err error) {
	agent = &logAgent{}
	agent.tailMan, err = tail.TailFile(file, tail.Config{
		ReOpen: true,
		Pipe:   false,
		Follow: true,
		Logger: nil,
	})
	if err != nil {
		return
	}
	//agent.tailMan.Logger =
	return
}

func (agent *logAgent) Graceful() (err error) {
	agent.offset, err = agent.tailMan.Tell()
	if err != nil {
		return
	}
	err = agent.tailMan.Stop()
	return
}

func (agent *logAgent) loop() {
	for {
		select {
		case line := <-agent.tailMan.Lines:
			if line.Err != nil {

			}
			agent.logSplit(line.Text)

		}
	}
}

func (agent *logAgent) logSplit(lineText string) (err error) {
	msgBody := strings.Fields(`2020-01-91 woami "test aaa"`)
	print(msgBody)
	return
}
