package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/AntonioMorales97/ecom-go/db/mock"
	db "github.com/AntonioMorales97/ecom-go/db/sqlc"
	"github.com/AntonioMorales97/ecom-go/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetProductOrderAPI(t *testing.T) {
	productOrder := randomProductOrder()

	testCases := []struct {
		name           string
		productOrderID int64
		buildStubs     func(store *mockdb.MockStore)
		checkResponse  func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:           "OK",
			productOrderID: productOrder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductOrder(gomock.Any(), gomock.Eq(productOrder.ID)).
					Times(1).
					Return(productOrder, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProductOrder(t, recorder.Body, productOrder)
			},
		},
		{
			name:           "NotFound",
			productOrderID: productOrder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductOrder(gomock.Any(), gomock.Eq(productOrder.ID)).
					Times(1).
					Return(db.ProductOrder{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:           "InternalServerError",
			productOrderID: productOrder.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductOrder(gomock.Any(), gomock.Eq(productOrder.ID)).
					Times(1).
					Return(db.ProductOrder{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:           "InvalidID",
			productOrderID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetProductOrder(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			//build stubs
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/order/%d", tc.productOrderID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateProductOrderAPI(t *testing.T) {
	productOrder := randomProductOrder()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"quantity":   productOrder.Quantity,
				"product_id": productOrder.ProductID,
				"owner":      productOrder.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProductOrderTx(gomock.Any(), gomock.Eq(db.CreateProductOrderParams{
					Quantity:  productOrder.Quantity,
					ProductID: productOrder.ProductID,
					Owner:     productOrder.Owner,
				})).
					Times(1).
					Return(productOrder, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchProductOrder(t, recorder.Body, productOrder)
			},
		},
		{
			name: "InvalidQuantity",
			body: gin.H{
				"quantity":   0,
				"product_id": productOrder.ProductID,
				"owner":      productOrder.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProductOrderTx(gomock.Any(), gomock.Any()).
					Times(0).
					Return(productOrder, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"quantity":   productOrder.Quantity,
				"product_id": productOrder.ProductID,
				"owner":      productOrder.Owner,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateProductOrderTx(gomock.Any(), gomock.Eq(db.CreateProductOrderParams{
					Quantity:  productOrder.Quantity,
					ProductID: productOrder.ProductID,
					Owner:     productOrder.Owner,
				})).
					Times(1).
					Return(db.ProductOrder{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/order"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomProductOrder() db.ProductOrder {
	return db.ProductOrder{
		ID:        util.RandomInt64(1, 1000),
		Quantity:  util.RandomInt32(1, 100),
		ProductID: util.RandomInt64(1, 1000),
		Owner:     util.RandomString(5),
	}
}

func requireBodyMatchProductOrder(t *testing.T, body *bytes.Buffer, productOrder db.ProductOrder) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotProductOrder db.ProductOrder
	err = json.Unmarshal(data, &gotProductOrder)
	require.NoError(t, err)
	require.Equal(t, productOrder, gotProductOrder)
}
