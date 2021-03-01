package main

type stationManager struct {
	queue   []train
	isBlock bool
}

func NewStationManager() *stationManager {
	return &stationManager{
		queue:   []train{},
		isBlock: false,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isBlock {
		s.queue = append(s.queue, t)
		return false
	}
	s.isBlock = true
	return true
}

func (s *stationManager) notifyDepart() {
	if s.isBlock == true {
		s.isBlock = false
	}
	if len(s.queue) > 0 {
		t := s.queue[0]
		t.permitArrival()
		s.queue = s.queue[1:]
	}
}
