package tasks_postgres

import core_postgres "github.com/abi-kan/golang-todoapp/internal/core/adapter/postgres"

type Pool struct {
	pool core_postgres.Pool
}

func NewPool(pool core_postgres.Pool) *Pool {
	return &Pool{pool}
}
