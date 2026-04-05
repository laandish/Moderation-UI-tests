package advertisements

type Statistics struct {
	Likes     int `json:"likes"`
	ViewCount int `json:"viewCount"`
	Contacts  int `json:"contacts"`
}

type CreateItemRequest struct {
	SellerID   int        `json:"sellerID"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	Statistics Statistics `json:"statistics"`
}

type CreateItemResponse struct {
	Status string `json:"status"`
}

type ItemResponse struct {
	ID         string     `json:"id"`
	SellerID   int        `json:"sellerId"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	Statistics Statistics `json:"statistics"`
	CreatedAt  string     `json:"createdAt"`
}

type ErrorResponse struct {
	Result any    `json:"result"`
	Status string `json:"status"`
}