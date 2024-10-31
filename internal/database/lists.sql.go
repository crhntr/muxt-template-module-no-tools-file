// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: lists.sql

package database

import (
	"context"
)

const deleteList = `-- name: DeleteList :exec
DELETE FROM lists WHERE id = $1
`

func (q *Queries) DeleteList(ctx context.Context, db DBTX, id int64) error {
	_, err := db.Exec(ctx, deleteList, id)
	return err
}

const insertList = `-- name: InsertList :one
INSERT INTO lists (name, description) VALUES ($1, $2) RETURNING id, name, description
`

type InsertListParams struct {
	Name        string
	Description string
}

func (q *Queries) InsertList(ctx context.Context, db DBTX, arg InsertListParams) (List, error) {
	row := db.QueryRow(ctx, insertList, arg.Name, arg.Description)
	var i List
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const selectList = `-- name: SelectList :one
SELECT id, name, description FROM lists WHERE id = $1
`

func (q *Queries) SelectList(ctx context.Context, db DBTX, id int64) (List, error) {
	row := db.QueryRow(ctx, selectList, id)
	var i List
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const selectLists = `-- name: SelectLists :many
SELECT id, name, description FROM lists
`

func (q *Queries) SelectLists(ctx context.Context, db DBTX) ([]List, error) {
	rows, err := db.Query(ctx, selectLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []List
	for rows.Next() {
		var i List
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateList = `-- name: UpdateList :one
UPDATE lists SET name = $2, description = $3 WHERE id = $1 RETURNING id, name, description
`

type UpdateListParams struct {
	ID          int64
	Name        string
	Description string
}

func (q *Queries) UpdateList(ctx context.Context, db DBTX, arg UpdateListParams) (List, error) {
	row := db.QueryRow(ctx, updateList, arg.ID, arg.Name, arg.Description)
	var i List
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
