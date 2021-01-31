package storage

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var (
	hostname,_ = os.Hostname()
)

func initDB(filename string) (s *sqlite, err error) {
	s = &sqlite{
		filename: filename,
	}
	err = s.Initialize()
	return
}

func Test_Watch(t *testing.T)  {
	s := sqlite{
		filename: "tmp.db",
	}
	err := s.Initialize()
	assert.NoError(t, err, fmt.Sprintf("fail: %v",err))
	select {
		case <- time.After(time.Hour) :

	}
}

func Test_Init(t *testing.T)  {
	s := sqlite{
		filename: "tmp.db",
	}
	err := s.Initialize()
	assert.NoError(t, err, fmt.Sprintf("fail: %v",err))
}

func Test_AddProject(t *testing.T)  {
	s,_ := initDB("tmp.db")
	assert.NotNil(t, s)
	err := s.AddProject(context.Background(), hostname, "test4", "test.log")
	assert.NoErrorf(t, err,"add project fail: %+v",err)
}


func Test_UpdateOffset(t *testing.T)  {
	s,_ := initDB("tmp.db")
	assert.NotNil(t, s)
	err := s.UpdateRead(context.Background(), "test2", 200, false)
	assert.NoErrorf(t, err,"update fail: %+v",err)
}

func Test_GetOffset(t *testing.T)  {
	s,_ := initDB("tmp.db")
	assert.NotNil(t, s)
	offset, err := s.Read(context.Background(), "test2")
	assert.NoError(t, err, fmt.Sprintf("update fail: %v",err))
	fmt.Printf("project [%s] offset: %d ", "test1", offset)
}