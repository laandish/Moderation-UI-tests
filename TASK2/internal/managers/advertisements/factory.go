package advertisements

import (
	"fmt"
	"math/rand"
	"time"

	clientAds "api-tests-template/internal/client/http/advertisements"
)

func init() {	
	rand.Seed(time.Now().UnixNano())
}

func RandomSellerID() int {
	return 111111 + rand.Intn(999999-111111+1)
}

func RandomName() string {
	return fmt.Sprintf("test-item-%d", time.Now().UnixNano())
}

func RandomPrice() int {
	return 1 + rand.Intn(100000)
}

func RandomStatistics() clientAds.Statistics {
	return clientAds.Statistics{
		Likes:     rand.Intn(1000),
		ViewCount: rand.Intn(10000),
		Contacts:  rand.Intn(500),
	}
}

func ValidCreateItemRequest() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       RandomName(),
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func ValidCreateItemRequestBySellerID(sellerID int) clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   sellerID,
		Name:       RandomName(),
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func ValidCreateItemRequestWithExactValues(
	sellerID int,
	name string,
	price int,
	statistics clientAds.Statistics,
) clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   sellerID,
		Name:       name,
		Price:      price,
		Statistics: statistics,
	}
}

func RequestWithMinSellerID() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   111111,
		Name:       RandomName(),
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithMaxSellerID() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   999999,
		Name:       RandomName(),
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithZeroStatistics() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID: RandomSellerID(),
		Name:     RandomName(),
		Price:    RandomPrice(),
		Statistics: clientAds.Statistics{
			Likes:     0,
			ViewCount: 0,
			Contacts:  0,
		},
	}
}

func RequestWithZeroPrice() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       RandomName(),
		Price:      0,
		Statistics: RandomStatistics(),
	}
}

func RequestWithNegativePrice() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       RandomName(),
		Price:      -100,
		Statistics: RandomStatistics(),
	}
}

func RequestWithNegativeStatistics() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID: RandomSellerID(),
		Name:     RandomName(),
		Price:    RandomPrice(),
		Statistics: clientAds.Statistics{
			Likes:     -1,
			ViewCount: -10,
			Contacts:  -2,
		},
	}
}

func RequestWithEmptyName() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       "",
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithLongName(length int) clientAds.CreateItemRequest {
	if length <= 0 {
		length = 300
	}

	runes := make([]rune, length)
	for i := range runes {
		runes[i] = 'a'
	}

	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       string(runes),
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithSpecialCharsInName() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       "!@#$%^&*()_+{}|:<>?[];',./`~",
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithScriptLikeName() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       "<script>alert(1)</script>",
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}

func RequestWithSQLLikeName() clientAds.CreateItemRequest {
	return clientAds.CreateItemRequest{
		SellerID:   RandomSellerID(),
		Name:       "' OR 1=1 --",
		Price:      RandomPrice(),
		Statistics: RandomStatistics(),
	}
}