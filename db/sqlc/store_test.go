package db

import (
	"context"
	"testing"

	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomProductTx(t *testing.T, createProductTxArg CreateProductTxParams, store Store) CreateProductResult {

	createProductResult, err := store.CreateProductTx(context.Background(), createProductTxArg)

	require.NoError(t, err)
	require.Equal(t, createProductTxArg.Quantity, createProductResult.Quantity)
	require.Equal(t, createProductTxArg.DescriptionLong, createProductResult.Product.DescriptionLong.String)

	//TODO: Add more validation + validation of created inventorys and stuff (? latter maybe not needed)

	return createProductResult
}

func getRandomCreateProductTxParams(t *testing.T, quantity int32) CreateProductTxParams {
	productType := createRandomProductType(t)
	productCategory := createRandomProductCategory(t)

	descriptionShort := util.RandomString(10)

	return CreateProductTxParams{
		Name:              util.RandomString(5),
		DescriptionLong:   "",
		DescriptionShort:  descriptionShort,
		Price:             util.RandomInt32(100, 1000),
		ProductTypeID:     productType.ID,
		ProductCategoryID: productCategory.ID,
		Quantity:          quantity,
	}
}

func TestCreateProductTx(t *testing.T) {
	store := NewStore(testDB)

	createProductTxParams := getRandomCreateProductTxParams(t, util.RandomInt32(1, 100))

	createRandomProductTx(t, createProductTxParams, store)
}

func TestCreateProductOrderTx(t *testing.T) {
	store := NewStore(testDB)

	createProductTxParams := getRandomCreateProductTxParams(t, 1000)
	productResult := createRandomProductTx(t, createProductTxParams, store)

	n := 5
	minAmount := int32(1)
	maxAmount := int32(199)

	errs := make(chan error)
	results := make(chan CreateProductOrderResult)

	// run n concurrent transfer transaction
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateProductOrderTx(context.Background(), CreateProductOrderParams{
				ProductID: productResult.Product.ID,
				Quantity:  util.RandomInt32(minAmount, maxAmount),
			})

			errs <- err
			results <- result
		}()
	}

	totalAmount := int32(0)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check product order
		productOrder := result.ProductOrder
		require.NotEmpty(t, productOrder)
		require.NotEmpty(t, productOrder.CreatedAt)
		require.NotEmpty(t, productOrder.UpdatedAt)
		require.Equal(t, productResult.Product.ID, productOrder.ProductID)
		require.True(t, productOrder.Quantity >= minAmount && productOrder.Quantity <= maxAmount)

		// increment total amount
		totalAmount += productOrder.Quantity
	}

	productInventory, err := store.GetProductInventoryForProduct(context.Background(), productResult.Product.ID)
	require.NoError(t, err)

	require.Equal(t, 1000-totalAmount, productInventory.Quantity)
}
