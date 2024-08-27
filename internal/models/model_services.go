package models

type Services struct {
	ServiceId string  `gorm:"primaryKey" json: "serviceId"`
	Name      string  `json:"serviceName"`
	Price     float32 `json:"price"`
}
