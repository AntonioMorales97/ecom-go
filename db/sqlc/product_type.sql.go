// Code generated by sqlc. DO NOT EDIT.
// source: product_type.sql

package db

import (
	"context"
)

const createProductType = `-- name: CreateProductType :one
INSERT INTO product_type (
    name
) VALUES (
    $1
) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateProductType(ctx context.Context, name string) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, createProductType, name)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductType = `-- name: DeleteProductType :exec
DELETE FROM product_type
WHERE id = $1
`

func (q *Queries) DeleteProductType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductType, id)
	return err
}

const getProductType = `-- name: GetProductType :one
SELECT id, name, created_at, updated_at FROM product_type
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductType(ctx context.Context, id int64) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, getProductType, id)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProductTypes = `-- name: ListProductTypes :many
SELECT id, name, created_at, updated_at FROM product_type
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductTypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProductTypes(ctx context.Context, arg ListProductTypesParams) ([]ProductType, error) {
	rows, err := q.db.QueryContext(ctx, listProductTypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductType{}
	for rows.Next() {
		var i ProductType
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductType = `-- name: UpdateProductType :one
UPDATE product_type
SET name = $2
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateProductTypeParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateProductType(ctx context.Context, arg UpdateProductTypeParams) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, updateProductType, arg.ID, arg.Name)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
