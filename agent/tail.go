package agent

import (
	"github.com/hpcloud/tail"
	"strings"
)

var (
	agent *logAgent
)

type logAgent struct {
	reader *tail.Tail
	offset int64
}

//NewTail new log tail instance
func NewTail(file string) (err error) {
	agent = &logAgent{}
	agent.reader, err = tail.TailFile(file, tail.Config{
		ReOpen: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 0,
		},
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
	agent.offset, err = agent.reader.Tell()
	if err != nil {
		return
	}
	err = agent.reader.Stop()
	return
}

func (agent *logAgent) loop() {
	for {
		select {
		case line := <-agent.reader.Lines:
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
