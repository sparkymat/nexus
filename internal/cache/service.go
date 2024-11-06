package cache

import (
	"time"
)

const ttl = time.Second * 60 * 5

type Entry struct {
	Value any
	Time  time.Time
}

func New() *Service {
	return &Service{
		store: map[string]Entry{},
	}
}

type Service struct {
	store map[string]Entry
}

func (s *Service) Read(key string) any {
	val, isOK := s.store[key]
	if !isOK || val.Time.Add(ttl).Before(time.Now()) {
		return nil
	}

	val.Time = time.Now()

	s.store[key] = val

	return val.Value
}

func (s *Service) Write(key string, value any) {
	s.store[key] = Entry{Value: value, Time: time.Now()}
}
