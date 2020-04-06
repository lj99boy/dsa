package tstruct

import "sync"

type Set struct {
	mux sync.Mutex
	h   map[string]interface{}
}

func (s Set) Insert(key string, val interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, ok := s.h[key]; ok {
		return
	}
	s.h[key] = val
}


