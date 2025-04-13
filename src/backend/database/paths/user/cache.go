package database_user

import (
	"context"
	"time"

	cache "github.com/Code-Hex/go-generics-cache"
)

type UserCache struct {
	userEmailToId *cache.Cache[string, string]
	userFromId    *cache.Cache[string, *UserProfile]
}

// setUser sets a user in the cache
func (c *UserCache) setUser(user *UserProfile) {
	c.userFromId.Set(user.ID, user, cache.WithExpiration(time.Second*30))
	c.userEmailToId.Set(user.Email, user.ID, cache.WithExpiration(time.Second*30))
}

// getUserByEmail gets a user from the cache by email
func (c *UserCache) getUserByEmail(email *string) *UserProfile {
	userId, found := c.userEmailToId.Get(*email)
	if found {
		return c.getUserById(&userId)
	}
	return nil
}

// getUserById gets a user from the cache by id
func (c *UserCache) getUserById(id *string) *UserProfile {
	user, found := c.userFromId.Get(*id)
	if found {
		return user
	}
	return nil
}

// NewUserCache creates a new user cache
func NewUserCache(ctx context.Context) *UserCache {
	userEmailToId := cache.NewContext[string, string](ctx)
	userFromId := cache.NewContext[string, *UserProfile](ctx)

	return &UserCache{
		userEmailToId: userEmailToId,
		userFromId:    userFromId,
	}
}
