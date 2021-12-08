package db

import (
	"context"
	"testing"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func getRandomProductType(t *testing.T) ProductType {
	productTypeID := util.RandomProductTypeID()
	productType, err := testQueries.GetProductType(context.Background(), productTypeID)
	require.NoError(t, err)
	require.NotEmpty(t, productType)

	require.Equal(t, productTypeID, productType.ID)
	require.NotEmpty(t, productType.Name)
	require.NotEmpty(t, productType.CreatedAt)
	require.NotEmpty(t, productType.UpdatedAt)
	return productType
}
func TestGetProductType(t *testing.T) {
	getRandomProductType(t)
}

func TestListProductTypes(t *testing.T) {

	arg := ListProductTypesParams{
		Limit:  5,
		Offset: 0,
	}

	productTypes, err := testQueries.ListProductTypes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, productTypes, 2)

	for _, productType := range productTypes {
		require.NotEmpty(t, productType)
	}
}
