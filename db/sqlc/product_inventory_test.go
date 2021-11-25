package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createTestProductInventory(t *testing.T, quantity int32) ProductInventory {
	productInventory, err := testQueries.CreateProductInventory(context.Background(), quantity)
	require.NoError(t, err)
	require.NotEmpty(t, productInventory)

	require.Equal(t, quantity, productInventory.Quantity)
	require.NotZero(t, productInventory.ID)
	require.NotEmpty(t, productInventory.CreatedAt)
	require.NotEmpty(t, productInventory.UpdatedAt)

	return productInventory
}

func TestCreateProductInventory(t *testing.T) {
	createTestProductInventory(t, util.RandomInt32(0, 100))
}

func TestGetProductInventory(t *testing.T) {
	productInventory1 := createTestProductInventory(t, util.RandomInt32(0, 100))
	productInventory2, err := testQueries.GetProductInventory(context.Background(), productInventory1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productInventory2)

	require.Equal(t, productInventory1.ID, productInventory2.ID)
	require.Equal(t, productInventory1.Quantity, productInventory2.Quantity)
	require.WithinDuration(t, productInventory1.CreatedAt, productInventory2.CreatedAt, time.Second)
	require.WithinDuration(t, productInventory1.UpdatedAt, productInventory2.UpdatedAt, time.Second)
}

func TestUpdateProductInventory(t *testing.T) {
	productInventory1 := createTestProductInventory(t, util.RandomInt32(0, 100))

	arg := UpdateProductInventoryQuantityParams{
		productInventory1.ID,
		util.RandomInt32(0, 10),
	}
	productInventory2, err := testQueries.UpdateProductInventoryQuantity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productInventory2)

	require.Equal(t, productInventory1.ID, productInventory2.ID)
	require.Equal(t, arg.Quantity, productInventory2.Quantity)
	require.WithinDuration(t, productInventory1.CreatedAt, productInventory2.CreatedAt, time.Second)
	require.NotEqual(t, productInventory1.UpdatedAt, productInventory2.UpdatedAt)
}

func TestDeleteProductInventory(t *testing.T) {
	productInventory1 := createTestProductInventory(t, util.RandomInt32(0, 100))
	err := testQueries.DeleteProductInventory(context.Background(), productInventory1.ID)
	require.NoError(t, err)

	productInventory2, err := testQueries.GetProductInventory(context.Background(), productInventory1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, productInventory2)
}

func TestListProductInventories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestProductInventory(t, util.RandomInt32(0, 100))
	}

	arg := ListProductInventoriesParams{
		Limit:  5,
		Offset: 5,
	}

	productInventories, err := testQueries.ListProductInventories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, productInventories, 5)

	for _, productType := range productInventories {
		require.NotEmpty(t, productType)
	}
}
