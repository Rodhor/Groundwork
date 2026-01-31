package database

import (
	. "Groundwork/backend/internal/domain"
	"context"
	"errors"
	"sync"
)

type UserDB struct {
	mu        sync.RWMutex
	Users     map[int64]User
	CurrentID int64
}

// Error messages
var (
	ErrUsernameExists = errors.New("username already exists")
	ErrUserNotFound   = errors.New("user not found")
)

func NewUserDB() *UserDB {
	return &UserDB{
		Users:     make(map[int64]User),
		CurrentID: 0,
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

	user.ID = db.CurrentID
	db.Users[user.ID] = *user
	db.CurrentID++
	return user, nil
}

func (db *UserDB) GetUserByID(ctx context.Context, id int64) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if user, ok := db.Users[id]; ok {
		return &user, nil
	}
	return nil, ErrUserNotFound
}

func (db *UserDB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	for _, user := range db.Users {
		if user.Username == username {
			return &user, nil
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

	db.Users[user.ID] = *user

	return nil
}

func (db *UserDB) DeleteUser(ctx context.Context, id int64) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.Users[id]; !ok {
		return ErrUserNotFound
	}

	delete(db.Users, id)

	return nil
}
