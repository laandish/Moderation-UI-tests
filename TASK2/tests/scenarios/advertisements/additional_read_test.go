package advertisements_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	factoryAds "api-tests-template/internal/factories/advertisements"
	managerAds "api-tests-template/internal/managers/advertisements"
)

const additionalBaseURL = "https://qa-internship.avito.com"

func TestGetStatisticByItemIDIsIdempotent(t *testing.T) {
	manager := managerAds.NewManager(additionalBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	first, firstStatus, err := manager.GetStatisticByItemID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, firstStatus)
	require.NotNil(t, first)
	require.NotEmpty(t, *first.Response)

	second, secondStatus, err := manager.GetStatisticByItemID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, secondStatus)
	require.NotNil(t, second)
	require.NotEmpty(t, *second.Response)

	firstStat := (*first.Response)[0]
	secondStat := (*second.Response)[0]

	require.Equal(t, firstStat.Likes, secondStat.Likes)
	require.Equal(t, firstStat.ViewCount, secondStat.ViewCount)
	require.Equal(t, firstStat.Contacts, secondStat.Contacts)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetItemsBySellerIDWithoutItems(t *testing.T) {
	manager := managerAds.NewManager(additionalBaseURL)

	// очень маловероятный sellerID, для которого мы не создавали объявления в тесте
	sellerID := 999998

	items, statusCode, err := manager.GetItemsBySellerID(sellerID)
	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, items)
	require.NotNil(t, items.Response)
}

func TestCreatedAtIsFilled(t *testing.T) {
	manager := managerAds.NewManager(additionalBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	got, getStatus, err := manager.GetItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, getStatus)
	require.NotNil(t, got)
	require.NotEmpty(t, *got.Response)

	firstItem := (*got.Response)[0]
	require.NotEmpty(t, firstItem.CreatedAt)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetItemsBySellerIDResponseTime(t *testing.T) {
	manager := managerAds.NewManager(additionalBaseURL)
	sellerID := factoryAds.RandomSellerID()

	request := factoryAds.ValidCreateItemRequestBySellerID(sellerID)
	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)

	start := time.Now()
	items, statusCode, err := manager.GetItemsBySellerID(sellerID)
	duration := time.Since(start)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, items)
	require.Less(t, duration.Seconds(), 2.0)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestGetItemByIDRepeatedSeriesStability(t *testing.T) {
	manager := managerAds.NewManager(additionalBaseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	for i := 0; i < 10; i++ {
		got, statusCode, err := manager.GetItemByID(created.ItemID)
		require.NoError(t, err)
		require.Equal(t, 200, statusCode)
		require.NotNil(t, got)
		require.NotEmpty(t, *got.Response)
	}

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}
