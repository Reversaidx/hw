package memorystorage

import (
	"errors"
	"sync"

	"github.com/Reversaidx/hw/hw12_13_14_15_calendar/internal/storage"
)

var ErrEventNotFound = errors.New("event is not found")

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
		id:   0,
	}
}

type Storage struct {
	data map[int]storage.Event
	mu   sync.RWMutex
	id   int
}

func (s *Storage) Add(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[s.id] = event
	s.id++
	return nil
}

func (s *Storage) Get(eventId int) (storage.Event, error) {
	if event, ok := s.data[eventId]; ok {
		return event, nil
	}
	return storage.Event{}, ErrEventNotFound
}

func (s *Storage) Change(eventId int, event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[eventId]; ok {
		s.data[eventId] = event
		return nil
	}
	return ErrEventNotFound
}

func (s *Storage) Delete(eventId int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, eventId)
	return nil
}

func (s *Storage) List() []storage.Event {
	eventList := make([]storage.Event, 0)
	for _, v := range s.data {
		eventList = append(eventList, v)
	}
	return eventList
}
