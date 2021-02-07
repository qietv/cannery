package storage

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"sync"
)

type kv struct {
	filename string
	conn     *bolt.DB
	locker   *sync.Mutex
}

func (k *kv) Initialize() (err error) {
	k.conn, err = bolt.Open(k.filename, 0600, &bolt.Options{
		Timeout:    0,
		NoGrowSync: false,
		ReadOnly:   false,
	})
	if err != nil {
		return
	}
	return
}
func (k *kv) AddProject(ctx context.Context, host, project, path string) (err error) {
	err = k.conn.Batch(func(tx *bolt.Tx) (err error) {
		var (
			bucket *bolt.Bucket
			info   []byte
			offset [8]byte
		)
		offset = [8]byte{}
		bucket, err = tx.CreateBucketIfNotExists([]byte(project))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		info, err = json.Marshal(&Project{
			Name:     project,
			Path:     path,
			Hostname: host,
			Status:   "PV",
		})
		if err != nil {
			return fmt.Errorf("add project: %s", err.Error())
		}
		err = bucket.Put([]byte("project"), info)
		if err != nil {
			return
		}
		binary.BigEndian.PutUint64(offset[0:], uint64(0))
		err = bucket.Put([]byte("offset"), offset[0:])
		return
	})
	return
}

func (k *kv) RemoveProject(ctx context.Context, project string) (err error) {
	err = k.conn.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(project))
	})
	return
}

func (k *kv) Read(ctx context.Context, project string) (offset int, err error) {
	err = k.conn.Update(func(tx *bolt.Tx) (err error) {
		var (
			bucket    *bolt.Bucket
			offsetBuf []byte
		)
		bucket = tx.Bucket([]byte(project))
		if bucket == nil {
			err = UnknownProject
			return
		}
		offsetBuf = bucket.Get([]byte("offset"))
		offset = int(binary.BigEndian.Uint64(offsetBuf))
		return
	})
	return
}

func (k *kv) UpdateRead(ctx context.Context, project string, offset int, replace bool) (err error) {
	err = k.conn.Update(func(tx *bolt.Tx) (err error) {
		var (
			bucket    *bolt.Bucket
			newOffset uint64
			toFill    []byte
		)
		toFill = make([]byte, 8)
		bucket = tx.Bucket([]byte(project))
		if bucket == nil {
			err = UnknownProject
			return
		}
		if replace {
			newOffset = uint64(offset) + binary.BigEndian.Uint64(bucket.Get([]byte("offset")))
		} else {
			newOffset = uint64(offset)
		}
		binary.BigEndian.PutUint64(toFill, newOffset)
		err = bucket.Put([]byte("offset"), toFill)
		if err != nil {
			return
		}
		err = tx.Bucket([]byte(project)).Put([]byte("offset"), toFill)
		return
	})
	return
}
