package advertisements_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	clientAds "api-tests-template/internal/client/http/advertisements"
	factoryAds "api-tests-template/internal/factories/advertisements"
	managerAds "api-tests-template/internal/managers/advertisements"
)

const finalBaseURL = "https://qa-internship.avito.com"

func TestGetItemsByInvalidNegativeSellerID(t *testing.T) {
	client := clientAds.NewClient(finalBaseURL)

	resp, err := client.GetItemsBySellerID(-1)
	require.NoError(t, err)
	require.Contains(t, []int{400, 404}, resp.StatusCode)
}

func TestDeleteOneOfTwoSellerItemsSecondStillExists(t *testing.T) {
	manager := managerAds.NewManager(finalBaseURL)
	sellerID := factoryAds.RandomSellerID()

	request1 := factoryAds.ValidCreateItemRequestBySellerID(sellerID)
	request2 := factoryAds.ValidCreateItemRequestBySellerID(sellerID)

	created1, status1, err1 := manager.CreateItem(request1)
	created2, status2, err2 := manager.CreateItem(request2)

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.Equal(t, 200, status1)
	require.Equal(t, 200, status2)
	require.NotEmpty(t, created1.ItemID)
	require.NotEmpty(t, created2.ItemID)

	_, deleteStatus, err := manager.DeleteItemByID(created1.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, deleteStatus)

	items, getStatus, err := manager.GetItemsBySellerID(sellerID)
	require.NoError(t, err)
	require.Equal(t, 200, getStatus)
	require.NotNil(t, items)
	require.NotEmpty(t, *items.Response)

	firstExists := false
	secondExists := false

	for _, item := range *items.Response {
		if item.ID == created1.ItemID {
			firstExists = true
		}
		if item.ID == created2.ItemID {
			secondExists = true
		}
	}

	require.False(t, firstExists, "удаленное объявление не должно присутствовать в списке")
	require.True(t, secondExists, "второе объявление должно остаться доступным")

	_, _, _ = manager.DeleteItemByID(created2.ItemID)
}

func TestGetItemByIDResponseContentType(t *testing.T) {
	manager := managerAds.NewManager(finalBaseURL)
	client := clientAds.NewClient(finalBaseURL)

	request := factoryAds.ValidCreateItemRequest()
	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	resp, err := client.GetItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Contains(t, resp.Headers.Get("Content-Type"), "application/json")

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetStatisticByItemIDResponseContentType(t *testing.T) {
	manager := managerAds.NewManager(finalBaseURL)
	client := clientAds.NewClient(finalBaseURL)

	request := factoryAds.ValidCreateItemRequest()
	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	resp, err := client.GetStatisticByItemID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Contains(t, resp.Headers.Get("Content-Type"), "application/json")

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetStatisticByItemIDRepeatedSeriesStability(t *testing.T) {
	manager := managerAds.NewManager(finalBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	var firstLikes, firstViews, firstContacts int

	for i := 0; i < 5; i++ {
		stat, statusCode, err := manager.GetStatisticByItemID(created.ItemID)
		require.NoError(t, err)
		require.Equal(t, 200, statusCode)
		require.NotNil(t, stat)
		require.NotEmpty(t, *stat.Response)

		current := (*stat.Response)[0]

		if i == 0 {
			firstLikes = current.Likes
			firstViews = current.ViewCount
			firstContacts = current.Contacts
		} else {
			require.Equal(t, firstLikes, current.Likes)
			require.Equal(t, firstViews, current.ViewCount)
			require.Equal(t, firstContacts, current.Contacts)
		}
	}

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}
