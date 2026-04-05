package advertisements

import (
	"fmt"
	"net/http"

	"api-tests-template/internal/client/http/common"
)

type Client struct {
	httpClient *common.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: common.NewClient(baseURL),
	}
}

// POST /api/1/item
func (c *Client) CreateItem(request CreateItemRequest) (*common.Response, error) {
	return c.httpClient.Do(
		http.MethodPost,
		"/api/1/item",
		request,
		nil,
	)
}

// GET /api/1/item/{id}
func (c *Client) GetItemByID(id string) (*common.Response, error) {
	path := fmt.Sprintf("/api/1/item/%s", id)
	return c.httpClient.Do(
		http.MethodGet,
		path,
		nil,
		nil,
	)
}

// GET /api/1/{sellerID}/item
func (c *Client) GetItemsBySellerID(sellerID int) (*common.Response, error) {
	path := fmt.Sprintf("/api/1/%d/item", sellerID)
	return c.httpClient.Do(
		http.MethodGet,
		path,
		nil,
		nil,
	)
}

// GET /api/1/statistic/{id}
func (c *Client) GetStatisticByItemID(id string) (*common.Response, error) {
	path := fmt.Sprintf("/api/1/statistic/%s", id)
	return c.httpClient.Do(
		http.MethodGet,
		path,
		nil,
		nil,
	)
}

// DELETE /api/2/item/{id}
func (c *Client) DeleteItemByID(id string) (*common.Response, error) {
	path := fmt.Sprintf("/api/2/item/%s", id)
	return c.httpClient.Do(
		http.MethodDelete,
		path,
		nil,
		nil,
	)
}