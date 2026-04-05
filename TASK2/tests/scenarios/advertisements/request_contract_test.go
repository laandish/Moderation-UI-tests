package advertisements_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/common"
)

const contractBaseURL = "https://qa-internship.avito.com"

func TestCreateItemWithStringSellerID(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	body := map[string]any{
		"sellerID": "123456",
		"name":     "string seller id",
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do(http.MethodPost, "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithStringPrice(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "string price",
		"price":    "1000",
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do(http.MethodPost, "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithInvalidStatisticsType(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	body := map[string]any{
		"sellerID":   123456,
		"name":       "invalid statistics type",
		"price":      1000,
		"statistics": "abc",
	}

	resp, err := client.Do(http.MethodPost, "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithExtraField(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "extra field item",
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
		"extraField": "unexpected",
	}

	resp, err := client.Do(http.MethodPost, "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Contains(t, []int{200, 400}, resp.StatusCode)
}

func TestCreateItemWithSellerIdInsteadOfSellerID(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	body := map[string]any{
		"sellerId": 123456,
		"name":     "wrong seller field",
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do(http.MethodPost, "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Contains(t, []int{200, 400}, resp.StatusCode)
}

func TestCreateItemWithInvalidJSON(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	raw := []byte(`{"sellerID":123456,"name":"broken json","price":1000,"statistics":{"likes":1,"viewCount":10,"contacts":1}`)
	req, err := http.NewRequest(http.MethodPost, contractBaseURL+"/api/1/item", bytes.NewBuffer(raw))
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.HTTPClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithoutContentType(t *testing.T) {
	client := common.NewClient(contractBaseURL)

	raw := []byte(`{"sellerID":123456,"name":"no content type","price":1000,"statistics":{"likes":1,"viewCount":10,"contacts":1}}`)
	req, err := http.NewRequest(http.MethodPost, contractBaseURL+"/api/1/item", bytes.NewBuffer(raw))
	require.NoError(t, err)

	req.Header.Set("Accept", "application/json")
	// Content-Type намеренно не ставим

	resp, err := client.HTTPClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Contains(t, []int{200, 400}, resp.StatusCode)
}
