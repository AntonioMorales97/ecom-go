-- name: CreateProductCategory :one
INSERT INTO product_category (
    name
) VALUES (
    $1
) RETURNING *;

-- name: GetProductCategory :one
SELECT * FROM product_category
WHERE id = $1 LIMIT 1;

-- name: ListProductCategories :many
SELECT * FROM product_category
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProductCategory :one
UPDATE product_category
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProductCategory :exec
DELETE FROM product_category
WHERE id = $1;