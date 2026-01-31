package database

import (
	. "Groundwork/backend/internal/domain"
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type UserDB struct {
	mu    sync.RWMutex
	Users map[uuid.UUID]*User
}

// Error messages
var (
	ErrUsernameExists = errors.New("username already exists")
	ErrUserNotFound   = errors.New("user not found")
)

func NewUserDB() *UserDB {
	return &UserDB{
		Users: make(map[uuid.UUID]*User),
	}
}

func (db *UserDB) AddNewUser(ctx context.Context, user *User) (*User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Closure for checking duplicate usernames
	checkDuplicateUsername := func(user *User) bool {
		for _, currentUsers := range db.Users {
			if currentUsers.Username == user.Username {
				return true
			}
		}
		return false
	}

	if checkDuplicateUsername(user) {
		return nil, ErrUsernameExists
	}

	id := uuid.New()
	user.ID = id
	db.Users[id] = user
	return user, nil
}

func (db *UserDB) GetUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if user, ok := db.Users[id]; ok {
		return user, nil
	}
	return nil, ErrUserNotFound
}

func (db *UserDB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, user := range db.Users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *UserDB) UpdateUser(ctx context.Context, user *User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Users[user.ID]; !ok {
		return ErrUserNotFound
	}

	db.Users[user.ID] = user

	return nil
}

func (db *UserDB) DeleteUser(ctx context.Context, id uuid.UUID) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Users[id]; !ok {
		return ErrUserNotFound
	}

	delete(db.Users, id)

	return nil
}
