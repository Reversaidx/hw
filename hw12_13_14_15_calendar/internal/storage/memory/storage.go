package memorystorage

import (
	"github.com/Reversaidx/hw/hw12_13_14_15_calendar/internal/storage"
	"sync"
)

type Store interface {
	Add(storage.Event) error
	Change(int, storage.Event) error
	Delete(int2 int) error
	List() []storage.Event
}

func New() Store {
	return &Storage{
		data: make(map[int]storage.Event),
		mu:   sync.RWMutex{},
	}
}

type Storage struct {
	data map[int]storage.Event
	mu   sync.RWMutex
}

func (s *Storage) Add(event storage.Event) error {
	s.mu.Lock()
	s.data[event.ID] = event
	s.mu.Unlock()
	return nil
}
func (s *Storage) Change(eventId int, event storage.Event) error {
	s.mu.Lock()
	s.data[eventId] = event
	s.mu.Unlock()
	return nil
}
func (s *Storage) Delete(eventId int) error {
	s.mu.Lock()
	delete(s.data, eventId)
	s.mu.Unlock()
	return nil
}
func (s *Storage) List() []storage.Event {
	eventList := make([]storage.Event, 0)
	for _, v := range s.data {
		eventList = append(eventList, v)
	}
	return eventList
}
