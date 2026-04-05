package advertisements_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	factoryAds "api-tests-template/internal/factories/advertisements"
	managerAds "api-tests-template/internal/managers/advertisements"
)

const baseURL = "https://qa-internship.avito.com"

func TestCreateItemPositive(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.ValidCreateItemRequest()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotNil(t, result.Response)
	require.NotEmpty(t, result.ItemID)
	require.NotEmpty(t, result.Response.Status)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateTwoSameItemsHaveDifferentIDs(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	sellerID := factoryAds.RandomSellerID()
	name := "same-name"
	price := 1000
	statistics := factoryAds.RandomStatistics()

	request1 := factoryAds.ValidCreateItemRequestWithExactValues(sellerID, name, price, statistics)
	request2 := factoryAds.ValidCreateItemRequestWithExactValues(sellerID, name, price, statistics)

	result1, statusCode1, err1 := manager.CreateItem(request1)
	result2, statusCode2, err2 := manager.CreateItem(request2)

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.Equal(t, 200, statusCode1)
	require.Equal(t, 200, statusCode2)

	require.NotNil(t, result1)
	require.NotNil(t, result2)
	require.NotEmpty(t, result1.ItemID)
	require.NotEmpty(t, result2.ItemID)
	require.NotEqual(t, result1.ItemID, result2.ItemID)

	_, _, _ = manager.DeleteItemByID(result1.ItemID)
	_, _, _ = manager.DeleteItemByID(result2.ItemID)
}

func TestCreateItemWithMinSellerID(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithMinSellerID()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	got, getStatus, err := manager.GetItemByID(result.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, getStatus)
	require.NotNil(t, got)
	require.NotEmpty(t, *got.Response)
	require.Equal(t, 111111, (*got.Response)[0].SellerID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateItemWithMaxSellerID(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithMaxSellerID()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	got, getStatus, err := manager.GetItemByID(result.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, getStatus)
	require.NotNil(t, got)
	require.NotEmpty(t, *got.Response)
	require.Equal(t, 999999, (*got.Response)[0].SellerID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateItemWithZeroStatistics(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithZeroStatistics()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 400, statusCode)
	require.Nil(t, result)
}

func TestCreateItemWithZeroPrice(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithZeroPrice()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 400, statusCode)
	require.Nil(t, result)
}

func TestCreateItemWithNegativePrice(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithNegativePrice()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 400, statusCode)
	require.Nil(t, result)
}

func TestCreateItemWithNegativeStatistics(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithNegativeStatistics()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 400, statusCode)
	require.Nil(t, result)
}

func TestCreateItemWithEmptyName(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithEmptyName()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 400, statusCode)
	require.Nil(t, result)
}

func TestCreateItemWithLongName(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithLongName(500)

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateItemWithSpecialCharsInName(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithSpecialCharsInName()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateItemWithScriptLikeName(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithScriptLikeName()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateItemWithSQLLikeName(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.RequestWithSQLLikeName()

	result, statusCode, err := manager.CreateItem(request)

	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, result)
	require.NotEmpty(t, result.ItemID)

	_, _, _ = manager.DeleteItemByID(result.ItemID)
}

func TestCreateGetDeleteItemE2E(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotNil(t, created)
	require.NotEmpty(t, created.ItemID)

	itemID := created.ItemID

	got, getStatus, err := manager.GetItemByID(itemID)
	require.NoError(t, err)
	require.Equal(t, 200, getStatus)
	require.NotNil(t, got)
	require.NotNil(t, got.Response)
	require.NotEmpty(t, *got.Response)

	firstItem := (*got.Response)[0]
	require.Equal(t, itemID, firstItem.ID)
	require.Equal(t, request.SellerID, firstItem.SellerID)
	require.Equal(t, request.Name, firstItem.Name)
	require.Equal(t, request.Price, firstItem.Price)

	deleted, deleteStatus, err := manager.DeleteItemByID(itemID)
	require.NoError(t, err)
	require.Equal(t, 200, deleteStatus)
	require.NotNil(t, deleted)
	require.Equal(t, 200, deleted.StatusCode)
}

func TestGetItemsBySellerIDPositive(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
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

	items, statusGet, err := manager.GetItemsBySellerID(sellerID)
	require.NoError(t, err)
	require.Equal(t, 200, statusGet)
	require.NotNil(t, items)
	require.NotNil(t, items.Response)
	require.NotEmpty(t, *items.Response)

	foundFirst := false
	foundSecond := false

	for _, item := range *items.Response {
		require.Equal(t, sellerID, item.SellerID)
		require.NotEmpty(t, item.ID)
		require.NotEmpty(t, item.Name)

		if item.ID == created1.ItemID {
			foundFirst = true
		}
		if item.ID == created2.ItemID {
			foundSecond = true
		}
	}

	require.True(t, foundFirst, "первое созданное объявление должно присутствовать в списке")
	require.True(t, foundSecond, "второе созданное объявление должно присутствовать в списке")

	_, _, _ = manager.DeleteItemByID(created1.ItemID)
	_, _, _ = manager.DeleteItemByID(created2.ItemID)
}

func TestGetStatisticByItemIDPositive(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	stat, statusCode, err := manager.GetStatisticByItemID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, statusCode)
	require.NotNil(t, stat)
	require.NotNil(t, stat.Response)
	require.NotEmpty(t, *stat.Response)

	firstStat := (*stat.Response)[0]
	require.Equal(t, request.Statistics.Likes, firstStat.Likes)
	require.Equal(t, request.Statistics.ViewCount, firstStat.ViewCount)
	require.Equal(t, request.Statistics.Contacts, firstStat.Contacts)

	_, _, _ = manager.DeleteItemByID(created.ItemID)
}

func TestDeleteItemTwice(t *testing.T) {
	manager := managerAds.NewManager(baseURL)
	request := factoryAds.ValidCreateItemRequest()

	created, createStatus, err := manager.CreateItem(request)
	require.NoError(t, err)
	require.Equal(t, 200, createStatus)
	require.NotEmpty(t, created.ItemID)

	_, firstDeleteStatus, err := manager.DeleteItemByID(created.ItemID)
	require.NoError(t, err)
	require.Equal(t, 200, firstDeleteStatus)

	_, secondDeleteStatus, err := manager.DeleteItemByID(created.ItemID)
	require.NoError(t, err)
	require.Contains(t, []int{400, 404}, secondDeleteStatus)
}