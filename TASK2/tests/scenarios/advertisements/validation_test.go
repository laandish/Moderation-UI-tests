package advertisements_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/common"
)

const validationBaseURL = "https://qa-internship.avito.com"

func TestCreateItemWithoutName(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do("POST", "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithoutSellerID(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"name":  "item without seller",
		"price": 1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do("POST", "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithoutPrice(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "item without price",
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do("POST", "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemWithoutStatistics(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "item without statistics",
		"price":    1000,
	}

	resp, err := client.Do("POST", "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 400, resp.StatusCode)
}

func TestCreateItemResponseContentType(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "content type test",
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	resp, err := client.Do("POST", "/api/1/item", body, nil)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Contains(t, resp.Headers.Get("Content-Type"), "application/json")
}

func TestCreateItemResponseTime(t *testing.T) {
	client := common.NewClient(validationBaseURL)

	body := map[string]any{
		"sellerID": 123456,
		"name":     "response-time-create",
		"price":    1000,
		"statistics": map[string]any{
			"likes":     1,
			"viewCount": 10,
			"contacts":  1,
		},
	}

	start := time.Now()
	resp, err := client.Do("POST", "/api/1/item", body, nil)
	duration := time.Since(start)

	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode)
	require.Less(t, duration.Seconds(), 2.0)
}
