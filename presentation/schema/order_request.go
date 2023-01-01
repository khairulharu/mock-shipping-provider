package schema

import "mock-shipping-provider/primitive"

type OrderRequest struct {
	Provider        string              `json:"provider"`
	Sender          primitive.Address   `json:"sender"`
	Recipient       primitive.Address   `json:"recipient"`
	Dimension       primitive.Dimension `json:"dimension"`
	Weight          float64             `json:"weight"`
	ItemDescription string              `json:"item_description"`
	ItemCategory    string              `json:"item_category"`
	Fragile         bool                `json:"fragile"`
}
