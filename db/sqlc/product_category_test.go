package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductCategory(t *testing.T) ProductCategory {
	name := util.RandomString(5)
	productCategory, err := testQueries.CreateProductCategory(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory)

	require.Equal(t, name, productCategory.Name)
	require.NotZero(t, productCategory.ID)
	require.NotEmpty(t, productCategory.CreatedAt)
	require.NotEmpty(t, productCategory.UpdatedAt)

	return productCategory
}

func TestCreateProductCategory(t *testing.T) {
	createRandomProductCategory(t)
}

func TestGetProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)
	productCategory2, err := testQueries.GetProductCategory(context.Background(), productCategory1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory2)

	require.Equal(t, productCategory1.ID, productCategory2.ID)
	require.Equal(t, productCategory1.Name, productCategory2.Name)
	require.WithinDuration(t, productCategory1.CreatedAt, productCategory2.CreatedAt, time.Second)
	require.WithinDuration(t, productCategory1.UpdatedAt, productCategory2.UpdatedAt, time.Second)
}

func TestUpdateProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)

	arg := UpdateProductCategoryParams{
		productCategory1.ID,
		util.RandomString(5),
	}
	productCategory2, err := testQueries.UpdateProductCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, productCategory2)

	require.Equal(t, productCategory1.ID, productCategory2.ID)
	require.Equal(t, arg.Name, productCategory2.Name)
	require.WithinDuration(t, productCategory1.CreatedAt, productCategory2.CreatedAt, time.Second)
	require.NotEqual(t, productCategory1.UpdatedAt, productCategory2.UpdatedAt)
}

func TestDeleteProductCategory(t *testing.T) {
	productCategory1 := createRandomProductCategory(t)
	err := testQueries.DeleteProductCategory(context.Background(), productCategory1.ID)
	require.NoError(t, err)

	productCategory2, err := testQueries.GetProductCategory(context.Background(), productCategory1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, productCategory2)
}

func TestListProductCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductCategory(t)
	}

	arg := ListProductCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	productCategories, err := testQueries.ListProductCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, productCategories, 5)

	for _, productType := range productCategories {
		require.NotEmpty(t, productType)
	}
}
