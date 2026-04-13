package entities

type StockItem struct {
	ID int64 `json:"id"`

	IdEquipment int64 `json:"id_equipment"`

	IdUnitStock int64 `json:"id_unit_stock"`

	StatusCode int64 `json:"status_code"`

	TypeMovement int64 `json:"type_movement"`

	Quantity int64 `json:"quantity"`
}
