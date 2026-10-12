package users_redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	core_redis "github.com/abi-kan/golang-todoapp/internal/core/adapter/redis"
	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *core_redis.Client
	ttl    time.Duration
}

func NewClient(client *core_redis.Client) *Client {
	return &Client{
		client: client,
		ttl:    10 * time.Minute,
	}
}

func (c *Client) SetUser(
	ctx context.Context,
	user domain.User,
) error {
	data, err := json.Marshal(userCacheModelFromDomain(user))
	if err != nil {
		return fmt.Errorf("encode user for redis: %w", err)
	}

	return c.client.Set(ctx, key(user.ID), data, c.ttl).Err()
}

func (c *Client) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	data, err := c.client.Get(ctx, key(id)).Bytes()
	if errors.Is(err, redis.Nil) {
		return domain.User{}, core_errors.ErrNotFound
	}

	if err != nil {
		return domain.User{}, fmt.Errorf("redis get user %d: %w", id, err)
	}

	var user userCacheModel
	if err := json.Unmarshal(data, &user); err != nil {
		return domain.User{}, fmt.Errorf("unmarshal cached user %d: %w", id, err)
	}

	return user.toDomain(), nil
}

func (c *Client) DeleteUser(ctx context.Context, id int) error {
	if err := c.client.Del(ctx, key(id)).Err(); err != nil {
		return fmt.Errorf("redis delete user %d: %w", id, err)
	}

	return nil
}

func key(id int) string {
	return fmt.Sprintf("user:%d", id)
}

type userCacheModel struct {
	ID          int     `json:"id"`
	Version     int     `json:"full_name"`
	FullName    string  `json:"phone_number"`
	PhoneNumber *string `json:"version"`
}

func userCacheModelFromDomain(user domain.User) userCacheModel {
	return userCacheModel{
		ID:          user.ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Version:     user.Version,
	}
}

func (m userCacheModel) toDomain() domain.User {
	return domain.User{
		ID:          m.ID,
		FullName:    m.FullName,
		PhoneNumber: m.PhoneNumber,
		Version:     m.Version,
	}
}
