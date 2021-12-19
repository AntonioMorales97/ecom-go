package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
)

type Store interface {
	Querier
	CreateProductTx(ctx context.Context, arg CreateProductTxParams) (CreateProductResult, error)
	CreateProductOrderTx(ctx context.Context, arg CreateProductOrderParams) (ProductOrder, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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
	ProductCategoryID *int64 `json:"product_category_id"`
	Quantity          int32  `json:"quantity"`
}

type CreateProductResult struct {
	Product  Product `json: "product"`
	Quantity int32   `json: "quantity"`
}

func (store *SQLStore) CreateProductTx(ctx context.Context, arg CreateProductTxParams) (CreateProductResult, error) {

	var result CreateProductResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		productInventory, err := q.CreateProductInventory(ctx, arg.Quantity)
		if err != nil {
			return err
		}

		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
			Name:               arg.Name,
			DescriptionLong:    util.ToNullString(&arg.DescriptionLong),
			DescriptionShort:   util.ToNullString(&arg.DescriptionShort),
			Price:              arg.Price,
			ProductTypeID:      arg.ProductTypeID,
			ProductCategoryID:  util.ToNullInt64(arg.ProductCategoryID),
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

func (store *SQLStore) CreateProductOrderTx(ctx context.Context, arg CreateProductOrderParams) (ProductOrder, error) {
	var result ProductOrder

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateProductOrder(ctx, CreateProductOrderParams{
			ProductID: arg.ProductID,
			Quantity:  arg.Quantity,
			Owner:     arg.Owner,
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
