-- name: CreateProductType :one
INSERT INTO product_type (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetProductType :one
SELECT * FROM product_type
WHERE id = $1 LIMIT 1;

-- name: ListProductTypes :many
SELECT * FROM product_type
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProductType :one
UPDATE product_type
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProductType :exec
DELETE FROM product_type
WHERE id = $1;