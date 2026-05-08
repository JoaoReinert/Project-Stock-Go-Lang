package entities

type Category struct {
	ID int64 `json:"id"`

	Name string `json:"name"`
}

type CategoryAndSubCategoryBalance struct {
	CategoryName    string `json:"category_name"`
	SubCategoryName string `json:"sub_category_name"`
	Balance         int64  `json:"balance"`
}

type CategoryBalance struct {
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
