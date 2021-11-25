-- name: CreateProductInventory :one
INSERT INTO product_inventory (
    quantity
) VALUES (
    $1
) RETURNING *;

-- name: GetProductInventory :one
SELECT * FROM product_inventory
WHERE id = $1 LIMIT 1;

-- name: GetProductInventoryForProduct :one
SELECT * FROM product_inventory
WHERE id = (SELECT product_inventory_id FROM product
            WHERE product.id = $1 LIMIT 1)
LIMIT 1;

-- name: GetProductInventoryForProductForUpdate :one
SELECT * FROM product_inventory
WHERE id = (SELECT product_inventory_id FROM product
            WHERE product.id = $1 LIMIT 1)
LIMIT 1 FOR NO KEY UPDATE;

-- name: ListProductInventories :many
SELECT * FROM product_inventory
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProductInventoryQuantity :one
UPDATE product_inventory
SET quantity = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProductInventoryQuantityForProduct :one
UPDATE product_inventory
SET quantity = quantity + $2
WHERE id = (SELECT product_inventory_id FROM product
            WHERE product.id = $1 LIMIT 1)
RETURNING *;

-- name: DeleteProductInventory :exec
DELETE FROM product_inventory
WHERE id = $1;