package models

type Products struct {
	ProductId string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"productName"`
	Price     float32 `json:"price"`
	VendorId  string  `json:"vendorId"`
}
