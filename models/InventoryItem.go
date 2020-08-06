package models

type InventoryItemInput struct {
	Cost    *float32 `json:"cost,omitempty"`
	Tracked bool     `json:"tracked,omitempty"`
}
