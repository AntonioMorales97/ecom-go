package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

//TODO: add validation?
type CreateProductTxParams struct {
	Name              string `json:"name"`
	DescriptionLong   string `json:"description_long"`
	DescriptionShort  string `json:"description_short"`
	Price             int32  `json:"price"`
	ProductTypeID     int64  `json:"product_type_id"`
	ProductCategoryID int64  `json:"product_category_id"`
	Quantity          int32  `json:"quantity"`
}

type CreateProductResult struct {
	Product  Product `json: "product"`
	Quantity int32   `json: "quantity"`
}

func (store *Store) CreateProductTx(ctx context.Context, arg CreateProductTxParams) (CreateProductResult, error) {

	var result CreateProductResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		productInventory, err := q.CreateProductInventory(ctx, arg.Quantity)
		if err != nil {
			return err
		}

		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
			Name:               arg.Name,
			DescriptionLong:    toNullString(&arg.DescriptionLong),
			DescriptionShort:   toNullString(&arg.DescriptionShort),
			Price:              arg.Price,
			ProductTypeID:      arg.ProductCategoryID,
			ProductCategoryID:  toNullInt64(&arg.ProductCategoryID),
			ProductInventoryID: productInventory.ID,
		})
		if err != nil {
			return err
		}

		result.Quantity = productInventory.Quantity

		return nil
	})

	return result, err

}

func toNullString(eventStr *string) sql.NullString {
	var nullString sql.NullString
	if len(*eventStr) == 0 {
		nullString.Valid = false
	} else {
		nullString.String = *eventStr
		nullString.Valid = true
	}
	return nullString
}

func toNullInt64(eventInt64 *int64) sql.NullInt64 {
	var nullInt64 sql.NullInt64
	nullInt64.Valid = (eventInt64 != nil)
	if !nullInt64.Valid {
		return nullInt64
	}

	nullInt64.Int64 = *eventInt64
	return nullInt64
}

type CreateProductOrderResult struct {
	ProductOrder ProductOrder `json: "product_order"`
}

func (store *Store) CreateProductOrderTx(ctx context.Context, arg CreateProductOrderParams) (CreateProductOrderResult, error) {
	var result CreateProductOrderResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.ProductOrder, err = q.CreateProductOrder(ctx, CreateProductOrderParams{
			ProductID: arg.ProductID,
			Quantity:  arg.Quantity,
		})
		if err != nil {
			return err
		}

		_, err = q.UpdateProductInventoryQuantityForProduct(ctx, UpdateProductInventoryQuantityForProductParams{
			ID:       arg.ProductID,
			Quantity: -arg.Quantity,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err

}
