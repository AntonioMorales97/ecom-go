-- name: CreateProductOrder :one
INSERT INTO product_order (
    quantity,
    product_id,
    owner
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetProductOrder :one
SELECT * FROM product_order
WHERE id = $1 LIMIT 1;

-- name: ListProductOrders :many
SELECT * FROM product_order
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateProductOrderQuantity :one
UPDATE product_order
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProductOrder :exec
DELETE FROM product_order
WHERE id = $1;