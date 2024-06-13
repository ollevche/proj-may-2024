package main

import (
	"fmt"
	"sort"
	"sync"
)

type Storage struct {
	m        sync.Mutex
	lastID   int
	allTrips map[int]Trip
	allUsers map[string]User
}

func NewStorage() *Storage {
	return &Storage{
		allTrips: make(map[int]Trip),
		allUsers: make(map[string]User),
	}
}

func (s *Storage) GetAllTrips() []Trip {
	s.m.Lock()
	defer s.m.Unlock()

	var trips = make([]Trip, 0, len(s.allTrips))

	for _, t := range s.allTrips {
		trips = append(trips, t)
	}

	sort.Slice(trips, func(i, j int) bool {
		return trips[i].ID < trips[j].ID
	})

	return trips
}

func (s *Storage) CreateOneTrip(t Trip) int {
	s.m.Lock()
	defer s.m.Unlock()

	fmt.Println("Trying to create trip")
	t.ID = s.lastID + 1
	s.allTrips[t.ID] = t
	s.lastID++
	fmt.Printf("Created trip. Last ID: %v\n", s.lastID)
	return t.ID
}

func (s *Storage) GetTripByID(id int) (Trip, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	t, ok := s.allTrips[id]
	return t, ok
}

func (s *Storage) DeleteTripByID(id int) bool {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.allTrips[id]
	if !ok {
		return false
	}

	delete(s.allTrips, id)
	return true
}

func (s *Storage) GetUserByUsername(username string) (User, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	u, ok := s.allUsers[username]
	return u, ok
}

func (s *Storage) CreateOneUser(u User) bool {
	s.m.Lock()
	defer s.m.Unlock()

	_, ok := s.allUsers[u.Username]
	if ok {
		return false
	}

	s.allUsers[u.Username] = u
	return true
}
