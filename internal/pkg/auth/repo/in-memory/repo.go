package in_memory

import (
	"errors"
	"ozonTech/internal/models"
	"sync"
)

type InMemoryAuthRepo struct {
	mu    sync.RWMutex
	Users map[int]*models.User
}

func NewInMemoryAuthRepo() *InMemoryAuthRepo {
	return &InMemoryAuthRepo{
		Users: make(map[int]*models.User),
	}
}

func (repo *InMemoryAuthRepo) CheckUser(name, password string) (int, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for id, user := range repo.Users {
		if user.Name == name && user.Password == password {
			return id, nil
		}
	}
	return 0, errors.New("user not found or incorrect password")

}

func (repo *InMemoryAuthRepo) CreateUser(name, password string) (int, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, u := range repo.Users {
		if u.Name == name {
			return 0, errors.New("user with this name already exists")
		}
	}

	id := len(repo.Users) + 1
	repo.Users[id] = &models.User{Name: name, Password: password, ID: id}
	return id, nil
}
