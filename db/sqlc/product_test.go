package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomCreateProductParams(t *testing.T) CreateProductParams {
	productType := createRandomProductType(t)
	productCategory := createRandomProductCategory(t)
	productInventory := createTestProductInventory(t, util.RandomInt32(0, 100))

	arg := CreateProductParams{
		util.RandomString(5),
		sql.NullString{String: "", Valid: false},
		sql.NullString{String: util.RandomString(10), Valid: true},
		util.RandomInt32(100, 1000),
		productType.ID,
		sql.NullInt64{Int64: productCategory.ID, Valid: true},
		productInventory.ID,
	}

	return arg
}

func validateCreatedProduct(t *testing.T, arg CreateProductParams, product Product) {
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.DescriptionLong, product.DescriptionLong)
	require.Equal(t, arg.DescriptionShort, product.DescriptionShort)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.ProductTypeID, product.ProductTypeID)
	require.Equal(t, arg.ProductCategoryID, product.ProductCategoryID)
	require.Equal(t, arg.ProductInventoryID, product.ProductInventoryID)
	require.NotEmpty(t, product.CreatedAt)
	require.NotEmpty(t, product.UpdatedAt)
}

func createRandomProduct(t *testing.T) Product {
	arg := createRandomCreateProductParams(t)

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	validateCreatedProduct(t, arg, product)

	return product
}

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.Name, product2.Name)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
	require.WithinDuration(t, product1.UpdatedAt, product2.UpdatedAt, time.Second)
}

func TestUpdateProduct(t *testing.T) {
	product1 := createRandomProduct(t)

	arg := UpdateProductNameParams{
		product1.ID,
		util.RandomString(5),
	}
	product2, err := testQueries.UpdateProductName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, arg.Name, product2.Name)
	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
	require.NotEqual(t, product1.UpdatedAt, product2.UpdatedAt)
}

func TestDeleteProduct(t *testing.T) {
	product1 := createRandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}

func TestListProduct(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
