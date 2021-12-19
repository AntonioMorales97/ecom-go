// Code generated by sqlc. DO NOT EDIT.
// source: product_order.sql

package db

import (
	"context"
)

const createProductOrder = `-- name: CreateProductOrder :one
INSERT INTO product_order (
    quantity,
    product_id,
    owner
) VALUES (
    $1, $2, $3
) RETURNING id, owner, quantity, product_id, created_at, updated_at
`

type CreateProductOrderParams struct {
	Quantity  int32  `json:"quantity"`
	ProductID int64  `json:"product_id"`
	Owner     string `json:"owner"`
}

func (q *Queries) CreateProductOrder(ctx context.Context, arg CreateProductOrderParams) (ProductOrder, error) {
	row := q.db.QueryRowContext(ctx, createProductOrder, arg.Quantity, arg.ProductID, arg.Owner)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Quantity,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProductOrder = `-- name: DeleteProductOrder :exec
DELETE FROM product_order
WHERE id = $1
`

func (q *Queries) DeleteProductOrder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductOrder, id)
	return err
}

const getProductOrder = `-- name: GetProductOrder :one
SELECT id, owner, quantity, product_id, created_at, updated_at FROM product_order
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductOrder(ctx context.Context, id int64) (ProductOrder, error) {
	row := q.db.QueryRowContext(ctx, getProductOrder, id)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Quantity,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProductOrders = `-- name: ListProductOrders :many
SELECT id, owner, quantity, product_id, created_at, updated_at FROM product_order
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProductOrders(ctx context.Context, arg ListProductOrdersParams) ([]ProductOrder, error) {
	rows, err := q.db.QueryContext(ctx, listProductOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductOrder{}
	for rows.Next() {
		var i ProductOrder
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Quantity,
			&i.ProductID,
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

const updateProductOrderQuantity = `-- name: UpdateProductOrderQuantity :one
UPDATE product_order
SET quantity = $2
WHERE id = $1
RETURNING id, owner, quantity, product_id, created_at, updated_at
`

type UpdateProductOrderQuantityParams struct {
	ID       int64 `json:"id"`
	Quantity int32 `json:"quantity"`
}

func (q *Queries) UpdateProductOrderQuantity(ctx context.Context, arg UpdateProductOrderQuantityParams) (ProductOrder, error) {
	row := q.db.QueryRowContext(ctx, updateProductOrderQuantity, arg.ID, arg.Quantity)
	var i ProductOrder
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Quantity,
		&i.ProductID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
