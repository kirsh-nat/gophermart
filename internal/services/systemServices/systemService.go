package systemservices

type Result struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual"`
}

type OrderRequest struct {
	Order string     `json:"order"`
	Good  []GoodDesc `json:"goods"`
}

type GoodDesc struct {
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
