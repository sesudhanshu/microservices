package models

type Vendor struct {
	VendorId string `gorm:"primaryKey" json:"vendorId"`
	Name     string `json:"vendorName"`
	Contact  string `json:"contact"`
	Phone    string `json:"phoneNumber"`
	Email    string `json:"emailAddress"`
	Address  string `json:"address"`
}
