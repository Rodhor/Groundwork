package database

import (
	"Groundwork/backend/internal/domain"
	"context"
	"errors"
	"sync"
)

type UserDB struct {
	mutext sync.RWMutex
	Users  map[int64]domain.User
}

func NewUserDB() *UserDB {
	return &UserDB{
		Users: make(map[int64]domain.User),
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
	db.mutext.Lock()
	defer db.mutext.Unlock()

	if db.CheckDuplicateUsername(ctx, user) {
		return &domain.User{}, errors.New("Username already exists")
	}

	db.Users[user.ID] = *user

	return user, nil
}

func (u *UserDB) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	u.mutext.RLock()
	defer u.mutext.RUnlock()

	if user, ok := u.Users[id]; ok {
		return &user, nil
	}
	return &domain.User{}, errors.New("User not found")
}

func (u *UserDB) UpdateUser(ctx context.Context, user *domain.User) error {
	u.mutext.Lock()
	defer u.mutext.Unlock()

	if _, ok := u.Users[user.ID]; !ok {
		return errors.New("User not found")
	}

	u.Users[user.ID] = *user

	return nil
}

func (u *UserDB) DeleteUser(ctx context.Context, id int64) error {
	u.mutext.Lock()
	defer u.mutext.Unlock()

	if _, ok := u.Users[id]; !ok {
		return errors.New("User not found")
	}

	delete(u.Users, id)

	return nil
}
