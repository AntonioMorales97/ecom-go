package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductOrder(t *testing.T) (ProductOrder, User) {
	product := createRandomProduct(t)
	user := createRandomUser(t)
	quantity := util.RandomInt32(1, 1000)

	productOrder, err := testQueries.CreateProductOrder(context.Background(), CreateProductOrderParams{
		ProductID: product.ID,
		Quantity:  quantity,
		Owner:     user.Username,
	})
	require.NoError(t, err)
	require.NotEmpty(t, productOrder)
	require.Equal(t, product.ID, productOrder.ProductID)
	require.Equal(t, quantity, productOrder.Quantity)
	require.Equal(t, productOrder.Owner, user.Username)
	return productOrder, user
}

func TestCreateProductOrder(t *testing.T) {
	createRandomProductOrder(t)
}

func TestGetProductOrder(t *testing.T) {
	productOrder1, _ := createRandomProductOrder(t)
	productOrder2, err := testQueries.GetProductOrder(context.Background(), productOrder1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productOrder2)

	require.Equal(t, productOrder1.ID, productOrder2.ID)
	require.Equal(t, productOrder1.ProductID, productOrder2.ProductID)
	require.Equal(t, productOrder1.Quantity, productOrder2.Quantity)
	require.Equal(t, productOrder1.Owner, productOrder2.Owner)
	require.WithinDuration(t, productOrder1.CreatedAt, productOrder2.CreatedAt, time.Second)
	require.WithinDuration(t, productOrder1.UpdatedAt, productOrder2.UpdatedAt, time.Second)
}

func TestUpdateProductOrder(t *testing.T) {
	productOrder1, _ := createRandomProductOrder(t)

	arg := UpdateProductOrderQuantityParams{
		productOrder1.ID,
		util.RandomInt32(productOrder1.Quantity+1, productOrder1.Quantity+1000),
	}
	productOrder2, err := testQueries.UpdateProductOrderQuantity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productOrder2)

	require.Equal(t, productOrder1.ID, productOrder2.ID)
	require.Equal(t, arg.Quantity, productOrder2.Quantity)
	require.NotEqual(t, productOrder1.Quantity, productOrder2.Quantity)
	require.WithinDuration(t, productOrder1.CreatedAt, productOrder2.CreatedAt, time.Second)
	require.NotEqual(t, productOrder1.UpdatedAt, productOrder2.UpdatedAt)
}

func TestDeleteProductOrder(t *testing.T) {
	productOrder1, _ := createRandomProductOrder(t)
	err := testQueries.DeleteProductOrder(context.Background(), productOrder1.ID)
	require.NoError(t, err)

	productOrder2, err := testQueries.GetProductOrder(context.Background(), productOrder1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, productOrder2)
}

func TestListProductOrder(t *testing.T) {
	var lastUser User
	for i := 0; i < 10; i++ {
		_, lastUser = createRandomProductOrder(t)
	}

	arg := ListProductOrdersParams{
		Owner:  lastUser.Username,
		Limit:  5,
		Offset: 0,
	}

	productOrders, err := testQueries.ListProductOrders(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productOrders)

	for _, productOrder := range productOrders {
		require.NotEmpty(t, productOrder)
		require.Equal(t, lastUser.Username, productOrder.Owner)
	}
}
