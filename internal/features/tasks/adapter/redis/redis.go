package tasks_redis

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

func (c *Client) SetTask(ctx context.Context, task domain.Task) error {
	data, err := json.Marshal(taskCacheModelFromDomain(task))
	if err != nil {
		return fmt.Errorf("encode task for redis: %w", err)
	}

	return c.client.Set(ctx, key(task.ID), data, c.ttl).Err()
}

func (c *Client) GetTask(ctx context.Context, id int) (domain.Task, error) {
	data, err := c.client.Get(ctx, key(id)).Bytes()
	if errors.Is(err, redis.Nil) {
		return domain.Task{}, core_errors.ErrNotFound
	}

	if err != nil {
		return domain.Task{}, fmt.Errorf("redis get task %d: %w", id, err)
	}

	var task taskCacheModel
	if err := json.Unmarshal(data, &task); err != nil {
		return domain.Task{}, fmt.Errorf("unmarshal cached task %d: %w", id, err)
	}

	return task.toDomain(), nil
}

func (c *Client) DeleteTask(ctx context.Context, id int) error {
	if err := c.client.Del(ctx, key(id)).Err(); err != nil {
		return fmt.Errorf("redis delete task %d: %w", id, err)
	}

	return nil
}

func key(id int) string {
	return fmt.Sprintf("task:%d", id)
}

type taskCacheModel struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorUserID int `json:"author_user_id"`
}

func taskCacheModelFromDomain(task domain.Task) taskCacheModel {
	return taskCacheModel{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func (t taskCacheModel) toDomain() domain.Task {
	return domain.Task{
		ID:           t.ID,
		Version:      t.Version,
		Title:        t.Title,
		Description:  t.Description,
		Completed:    t.Completed,
		CreatedAt:    t.CreatedAt,
		CompletedAt:  t.CompletedAt,
		AuthorUserID: t.AuthorUserID,
	}
}
