package entities

type SubCategory struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	IdCategory int64 `json:"id_category"`
}
