package advertisements_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	clientAds "api-tests-template/internal/client/http/advertisements"
	factoryAds "api-tests-template/internal/factories/advertisements"
	managerAds "api-tests-template/internal/managers/advertisements"
)

const readDeleteBaseURL = "https://qa-internship.avito.com"

func TestGetItemByNonExistingID(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.GetItemByID("00000000-0000-0000-0000-000000000000")
	require.NoError(t, err)
	require.Equal(t, 404, resp.StatusCode)
}

func TestGetStatisticByNonExistingID(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.GetStatisticByItemID("00000000-0000-0000-0000-000000000000")
	require.NoError(t, err)
	require.Equal(t, 404, resp.StatusCode)
}

func TestDeleteNonExistingID(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.DeleteItemByID("00000000-0000-0000-0000-000000000000")
	require.NoError(t, err)
	require.Equal(t, 404, resp.StatusCode)
}

func TestGetItemByInvalidIDFormat(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.GetItemByID("!!!invalid!!!")
	require.NoError(t, err)
	require.Contains(t, []int{400, 404}, resp.StatusCode)
}

func TestGetStatisticByInvalidIDFormat(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.GetStatisticByItemID("!!!invalid!!!")
	require.NoError(t, err)
	require.Contains(t, []int{400, 404}, resp.StatusCode)
}

func TestDeleteByInvalidIDFormat(t *testing.T) {
	client := clientAds.NewClient(readDeleteBaseURL)

	resp, err := client.DeleteItemByID("!!!invalid!!!")
	require.NoError(t, err)
	require.Contains(t, []int{400, 404}, resp.StatusCode)
}

func TestGetItemByIDIsIdempotent(t *testing.T) {
	manager := managerAds.NewManager(readDeleteBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	first, firstStatus, err := manager.GetItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, firstStatus)
	require.NotNil(t, first)
	require.NotEmpty(t, *first.Response)

	second, secondStatus, err := manager.GetItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, secondStatus)
	require.NotNil(t, second)
	require.NotEmpty(t, *second.Response)

	firstItem := (*first.Response)[0]
	secondItem := (*second.Response)[0]

	require.Equal(t, firstItem.ID, secondItem.ID)
	require.Equal(t, firstItem.SellerID, secondItem.SellerID)
	require.Equal(t, firstItem.Name, secondItem.Name)
	require.Equal(t, firstItem.Price, secondItem.Price)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetItemsBySellerIDIsIdempotent(t *testing.T) {
	manager := managerAds.NewManager(readDeleteBaseURL)
	sellerID := factoryAds.RandomSellerID()

	request1 := factoryAds.ValidCreateItemRequestBySellerID(sellerID)
	request2 := factoryAds.ValidCreateItemRequestBySellerID(sellerID)

	created1, status1, err1 := manager.CreateItem(request1)
	created2, status2, err2 := manager.CreateItem(request2)

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.Equal(t, 200, status1)
	require.Equal(t, 200, status2)

	first, firstStatus, err := manager.GetItemsBySellerID(sellerID)
	require.NoError(t, err)
	require.Equal(t, 200, firstStatus)
	require.NotNil(t, first)
	require.NotEmpty(t, *first.Response)

	second, secondStatus, err := manager.GetItemsBySellerID(sellerID)
	require.NoError(t, err)
	require.Equal(t, 200, secondStatus)
	require.NotNil(t, second)
	require.NotEmpty(t, *second.Response)

	require.Equal(t, len(*first.Response), len(*second.Response))

	_, _, _ = manager.DeleteItemByID(created1.ItemID)
	_, _, _ = manager.DeleteItemByID(created2.ItemID)
}

func TestGetDeletedItemReturnsNotFound(t *testing.T) {
	manager := managerAds.NewManager(readDeleteBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	_, deleteStatus, err := manager.DeleteItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, deleteStatus)

	client := clientAds.NewClient(readDeleteBaseURL)
	resp, err := client.GetItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 404, resp.StatusCode)
}

func TestGetItemByIDResponseTime(t *testing.T) {
	manager := managerAds.NewManager(readDeleteBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)

	client := clientAds.NewClient(readDeleteBaseURL)

	start := time.Now()
	resp, err := client.GetItemByID(created.ItemID)
	duration := time.Since(start)

	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Less(t, duration.Seconds(), 2.0)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}
