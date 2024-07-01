package user

import (
	"sync"
)

type InMemStorage struct {
	usersM sync.Mutex
	users  []User
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{}
}

func (s *InMemStorage) Create(u User) {
	s.usersM.Lock()
	defer s.usersM.Unlock()

	s.users = append(s.users, u)
}
