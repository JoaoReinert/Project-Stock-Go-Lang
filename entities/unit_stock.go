package entities

type UnitStock struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	StatusCode int64 `json:"status_code"`
}

type UnitStockBalance struct {
	Balance int64  `json:"balance"`
	Name    string `json:"name"`
}

type CategoryBalance struct {
	CategoryName    string `json:"category_name"`
	SubCategoryName string `json:"sub_category_name"`
	Balance         int64  `json:"balance"`
}
