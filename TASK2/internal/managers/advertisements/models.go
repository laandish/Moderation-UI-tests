package advertisements

import clientAds "api-tests-template/internal/client/http/advertisements"

type CreateItemResult struct {
	Response *clientAds.CreateItemResponse
	ItemID   string
}

type GetItemResult struct {
	Response *[]clientAds.ItemResponse
}

type GetItemsBySellerResult struct {
	Response *[]clientAds.ItemResponse
}

type GetStatisticResult struct {
	Response *[]clientAds.Statistics
}

type DeleteItemResult struct {
	StatusCode int
}