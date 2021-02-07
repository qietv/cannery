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
	hostname, _ = os.Hostname()
)

func initDB(filename string) (s *sqlite, err error) {
	s = &sqlite{
		filename: filename,
	}
	err = s.Initialize()
	return
}
func initKV(filename string) (s *kv, err error) {
	s = &kv{
		filename: filename,
	}
	err = s.Initialize()
	return
}

func Test_Watch(t *testing.T) {
	s := sqlite{
		filename: "tmp.db",
	}
	err := s.Initialize()
	assert.NoError(t, err, fmt.Sprintf("fail: %v", err))
	select {
	case <-time.After(time.Hour):

	}
}

func Test_Init(t *testing.T) {
	s := sqlite{
		filename: "tmp.db",
	}
	err := s.Initialize()
	assert.NoError(t, err, fmt.Sprintf("fail: %v", err))
}

func Test_AddProject(t *testing.T) {
	s, _ := initDB("tmp.db")
	assert.NotNil(t, s)
	err := s.AddProject(context.Background(), hostname, "test4", "test.log")
	assert.NoErrorf(t, err, "add project fail: %+v", err)
}

func Test_UpdateOffset(t *testing.T) {
	s, _ := initDB("tmp.db")
	assert.NotNil(t, s)
	err := s.UpdateRead(context.Background(), "test2", 200, false)
	assert.NoErrorf(t, err, "update fail: %+v", err)
}

func Test_GetOffset(t *testing.T) {
	s, _ := initDB("tmp.db")
	assert.NotNil(t, s)
	offset, err := s.Read(context.Background(), "test2")
	assert.NoError(t, err, fmt.Sprintf("get fail: %v", err))
	fmt.Printf("project [%s] offset: %d ", "test1", offset)
}

func Test_AddProjectKV(t *testing.T) {
	s, _ := initKV("kv.db")
	assert.NotNil(t, s)
	err := s.AddProject(context.Background(), hostname, "test2", "test.log")
	assert.NoErrorf(t, err, "add project fail: %+v", err)
}

func Test_UpOffsetKV(t *testing.T) {
	k, _ := initKV("kv.db")
	println(k.conn)
	assert.NotNil(t, k)
	err := k.UpdateRead(context.Background(), "test2", 1000, true)
	assert.NoError(t, err, fmt.Sprintf("update fail: %v", err))
}

func Test_OffsetKV(t *testing.T) {
	k, _ := initKV("kv.db")
	println(k.conn)
	assert.NotNil(t, k)
	offset, err := k.Read(context.Background(), "test2")
	assert.NoError(t, err, fmt.Sprintf("get fail: %v", err))
	fmt.Printf("project [%s] offset: %d ", "test2", offset)
}
