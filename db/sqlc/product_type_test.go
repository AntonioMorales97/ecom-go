package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductType(t *testing.T) ProductType {
	name := util.RandomString(5)
	productType, err := testQueries.CreateProductType(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, productType)

	require.Equal(t, name, productType.Name)
	require.NotZero(t, productType.ID)
	require.NotEmpty(t, productType.CreatedAt)
	require.NotEmpty(t, productType.UpdatedAt)

	return productType
}

func TestCreateProductType(t *testing.T) {
	createRandomProductType(t)
}

func TestGetProductType(t *testing.T) {
	productType1 := createRandomProductType(t)
	productType2, err := testQueries.GetProductType(context.Background(), productType1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productType2)

	require.Equal(t, productType1.ID, productType2.ID)
	require.Equal(t, productType1.Name, productType2.Name)
	require.WithinDuration(t, productType1.CreatedAt, productType2.CreatedAt, time.Second)
	require.WithinDuration(t, productType1.UpdatedAt, productType2.UpdatedAt, time.Second)
}

func TestUpdateProductType(t *testing.T) {
	productType1 := createRandomProductType(t)

	arg := UpdateProductTypeParams{
		productType1.ID,
		util.RandomString(5),
	}
	productType2, err := testQueries.UpdateProductType(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productType2)

	require.Equal(t, productType1.ID, productType2.ID)
	require.Equal(t, arg.Name, productType2.Name)
	require.WithinDuration(t, productType1.CreatedAt, productType2.CreatedAt, time.Second)
	require.NotEqual(t, productType1.UpdatedAt, productType2.UpdatedAt)
}

func TestDeleteProductType(t *testing.T) {
	productType1 := createRandomProductType(t)
	err := testQueries.DeleteProductType(context.Background(), productType1.ID)
	require.NoError(t, err)

	productType2, err := testQueries.GetProductType(context.Background(), productType1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, productType2)
}

func TestListProductTypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductType(t)
	}

	arg := ListProductTypesParams{
		Limit:  5,
		Offset: 5,
	}

	productTypes, err := testQueries.ListProductTypes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, productTypes, 5)

	for _, productType := range productTypes {
		require.NotEmpty(t, productType)
	}
}
