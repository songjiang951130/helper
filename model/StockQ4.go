package model

type StockQ4 struct {
	ID     uint   `gorm:"primaryKey"`
	Symbol string `gorm:"type:char(16);not null;default:;index:ip_idx"`
	Name   string
}
