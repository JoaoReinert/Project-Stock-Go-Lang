package entities

type Equipment struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	IdSubCategory int64 `json:"id_sub_category"`

	StatusCode int64 `json:"status_code"`
}
