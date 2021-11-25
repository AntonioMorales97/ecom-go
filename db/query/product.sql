-- name: CreateProduct :one
INSERT INTO product (
    name,
    description_long,
    description_short,
    price,
    product_type_id,
    product_category_id,
    product_inventory_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProductName :one
UPDATE product
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;