
-- name: GetProductType :one
SELECT * FROM product_type
WHERE id = $1 LIMIT 1;

-- name: ListProductTypes :many
SELECT * FROM product_type
ORDER BY id
LIMIT $1
OFFSET $2;