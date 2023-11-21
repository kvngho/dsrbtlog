package server

import (
	"fmt"
	"sync"
)

// Log 는 뮤텍스를 이용해, 잠금, 읽기 잠금을 수행할 수 있고 record를 이용해 로그를 기록한다
type Log struct {
	mu      sync.RWMutex
	records []Record
}

// NewLog : 새로운 로그 생성 함수
func NewLog() *Log {
	return &Log{}
}

// Append : 로그를 기존 로그에 추가한다
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// 기존 로그를 읽는다
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  string `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not Found!")
