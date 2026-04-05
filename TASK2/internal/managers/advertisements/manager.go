package advertisements

import (
	"fmt"

	clientAds "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/client/http/common"
)

type Manager struct {
	client *clientAds.Client
}

func NewManager(baseURL string) *Manager {
	return &Manager{
		client: clientAds.NewClient(baseURL),
	}
}

func (m *Manager) CreateItem(request clientAds.CreateItemRequest) (*CreateItemResult, int, error) {
	resp, err := m.client.CreateItem(request)
	if err != nil {
		return nil, 0, fmt.Errorf("create item request failed: %w", err)
	}

	decoded, err := common.Decode[clientAds.ItemResponse](resp)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to decode create item response: %w", err)
	}

	return &CreateItemResult{
		Response: decoded,
	}, resp.StatusCode, nil
}

func (m *Manager) GetItemByID(id string) (*GetItemResult, int, error) {
	resp, err := m.client.GetItemByID(id)
	if err != nil {
		return nil, 0, fmt.Errorf("get item by id request failed: %w", err)
	}

	decoded, err := common.Decode[[]clientAds.ItemResponse](resp)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to decode get item by id response: %w", err)
	}

	return &GetItemResult{
		Response: decoded,
	}, resp.StatusCode, nil
}

func (m *Manager) GetItemsBySellerID(sellerID int) (*GetItemsBySellerResult, int, error) {
	resp, err := m.client.GetItemsBySellerID(sellerID)
	if err != nil {
		return nil, 0, fmt.Errorf("get items by seller id request failed: %w", err)
	}

	decoded, err := common.Decode[[]clientAds.ItemResponse](resp)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to decode get items by seller id response: %w", err)
	}

	return &GetItemsBySellerResult{
		Response: decoded,
	}, resp.StatusCode, nil
}

func (m *Manager) GetStatisticByItemID(id string) (*GetStatisticResult, int, error) {
	resp, err := m.client.GetStatisticByItemID(id)
	if err != nil {
		return nil, 0, fmt.Errorf("get statistic by item id request failed: %w", err)
	}

	decoded, err := common.Decode[[]clientAds.Statistics](resp)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to decode get statistic response: %w", err)
	}

	return &GetStatisticResult{
		Response: decoded,
	}, resp.StatusCode, nil
}

func (m *Manager) DeleteItemByID(id string) (*DeleteItemResult, int, error) {
	resp, err := m.client.DeleteItemByID(id)
	if err != nil {
		return nil, 0, fmt.Errorf("delete item request failed: %w", err)
	}

	return &DeleteItemResult{
		StatusCode: resp.StatusCode,
	}, resp.StatusCode, nil
}