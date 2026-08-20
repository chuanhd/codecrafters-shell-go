package infra

import (
	"fmt"
)

type CompletionsStore interface {
	Set(cmd string, script string)
	Get(cmd string) (string, error)
}

type InMemoryCompletionsStore struct {
	mapCmdScrpt map[string]string
}

func NewInMemoryCompletionsStore() *InMemoryCompletionsStore {
	return &InMemoryCompletionsStore{
		mapCmdScrpt: make(map[string]string),
	}
}

func (s *InMemoryCompletionsStore) Set(cmd string, script string) {
	s.mapCmdScrpt[cmd] = script
	// fmt.Printf("check %s \n", s.mapCmdScrpt[cmd])
}

func (s *InMemoryCompletionsStore) Get(cmd string) (string, error) {
	val, ok := s.mapCmdScrpt[cmd]
	if ok {
		return val, nil
	}
	return "", fmt.Errorf("complete: %s: no completion specification", cmd)
}
