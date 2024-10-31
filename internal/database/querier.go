// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type Querier interface {
	DeleteList(ctx context.Context, db DBTX, id int64) error
	DeleteTask(ctx context.Context, db DBTX, id int64) error
	InsertList(ctx context.Context, db DBTX, arg InsertListParams) (List, error)
	InsertTask(ctx context.Context, db DBTX, arg InsertTaskParams) (Task, error)
	SelectList(ctx context.Context, db DBTX, id int64) (List, error)
	SelectLists(ctx context.Context, db DBTX) ([]List, error)
	SelectTasksForList(ctx context.Context, db DBTX, listID int64) ([]Task, error)
	UpdateList(ctx context.Context, db DBTX, arg UpdateListParams) (List, error)
	UpdateTaskSetCompletedAtNow(ctx context.Context, db DBTX, id int64) (Task, error)
}

var _ Querier = (*Queries)(nil)
