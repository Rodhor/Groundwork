package database

import (
	"Groundwork/backend/internal/domain"
	"context"
	"errors"
	"sync"
)

type UserDB struct {
	mu        sync.RWMutex
	Users     map[int64]domain.User
	CurrentID int64
}

func NewUserDB() *UserDB {
	return &UserDB{
		Users:     make(map[int64]domain.User),
		CurrentID: 0,
	}
}

func (db *UserDB) CheckDuplicateUsername(ctx context.Context, user *domain.User) bool {
	for _, currentUsers := range db.Users {
		if currentUsers.Username == user.Username {
			return true
		}
	}
	return false
}

func (db *UserDB) AddNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.CheckDuplicateUsername(ctx, user) {
		return &domain.User{}, errors.New("username already exists")
	}

	user.ID = db.CurrentID
	db.Users[user.ID] = *user

	db.CurrentID++
	return user, nil
}

func (db *UserDB) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if user, ok := db.Users[id]; ok {
		return &user, nil
	}
	return &domain.User{}, errors.New("User not found")
}

func (db *UserDB) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, user := range db.Users {
		if user.Username == username {
			return &user, nil
		}
	}
	return &domain.User{}, errors.New("User not found")
}

func (db *UserDB) UpdateUser(ctx context.Context, user *domain.User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Users[user.ID]; !ok {
		return errors.New("User not found")
	}

	db.Users[user.ID] = *user

	return nil
}

func (db *UserDB) DeleteUser(ctx context.Context, id int64) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Users[id]; !ok {
		return errors.New("User not found")
	}

	delete(db.Users, id)

	return nil
}
