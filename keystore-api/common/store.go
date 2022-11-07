package common

import (
	"sync"
)

type Store struct {
	sync.RWMutex
	Data map[string]string
}

func (s *Store) PutValue(key string, value string) error {
	// lock store to prevent concurrent updates
	s.Lock()
	s.Data[key] = value
	s.Unlock()

	return nil
}

func (s *Store) GetValue(key string) (string, error) {
	s.RLock()
	value, exists := s.Data[key]
	s.RUnlock()

	if !exists {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func (s *Store) DeleteValue(key string) error {
	s.Lock()
	delete(s.Data, key)
	s.Unlock()

	return nil
}
