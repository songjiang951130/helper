package model

import (
	"fmt"
	"time"
)

type StockQ4 struct {
	ID         uint32  `gorm:"primaryKey"`
	Symbol     string  `gorm:"type:char(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '代码';index:symbol_idx"`
	Name       string  `gorm:"type:varchar(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '名字'"`
	ReportName string  `gorm:"type:varchar(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '报告期名字'"`
	ReportTime uint    `gorm:"type:int unsigned NOT NULL DEFAULT '0' COMMENT '报告期'"`
	Roe        float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '净资产收益率'"`
	NetSell    float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '净利率'"`
	GrossSell  float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '毛利率'"`
	AssertRate float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '资产负债率'"`
	Type       string  `gorm:"type:char(2) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT ''"`
}

type North struct {
	ID         int64     `gorm:"primaryKey"`
	Symbol     string    `gorm:"column:symbol" json:"symbol" form:"symbol"`
	ReportRate float64   `gorm:"column:report_rate" json:"report_rate" form:"report_rate"`
	Rate       float64   `gorm:"column:rate" json:"rate" form:"rate"`
	UpdateAt   time.Time `gorm:"column:update_at" json:"update_at" form:"update_at"`
}

type Stock struct {
	ID     int64   `gorm:"primaryKey"`
	Symbol string  `gorm:"column:symbol" json:"symbol" form:"symbol"`
	Name   string  `gorm:"column:name" json:"name" form:"name"`
	Count  int64   `gorm:"column:count" json:"count" form:"count"`
	Price  float64 `gorm:"column:price" json:"price" form:"price"`
	State  int64   `gorm:"column:state" json:"state" form:"state"`
}

type StockDaily struct {
	ID        int64   `gorm:"primaryKey"`
	Symbol    string  `gorm:"column:symbol" json:"symbol" form:"symbol"`
	Name      string  `gorm:"column:name" json:"name" form:"name"`
	Timestamp int64   `gorm:"column:timestamp" json:"timestamp" form:"timestamp"`
	High      float64 `gorm:"column:high" json:"high" form:"high"`
	Low       float64 `gorm:"column:low" json:"low" form:"low"`
	Close     float64 `gorm:"column:close" json:"close" form:"close"`
	Open      float64 `gorm:"column:open" json:"open" form:"open"`
	Pe        float64 `gorm:"column:pe" json:"pe" form:"pe"`
	Pb        float64 `gorm:"column:pb" json:"pb" form:"pb"`
	Volume    int64   `gorm:"column:volume" json:"volume" form:"volume"`
}

func (StockDaily) TableName() string {
	return "stock_daily"
}

type StockSubscribe struct {
	ID             int64   `gorm:"primaryKey"`
	Symbol         string  `gorm:"column:symbol" json:"symbol" form:"symbol"`
	Name           string  `gorm:"column:name" json:"name" form:"name"`
	Price          float64 `gorm:"column:price" json:"price" form:"price"`
	NoticeLowPrice float64 `gorm:"column:notice_low_price" json:"notice_low_price" form:"notice_low_price"`
	OverPrice      float64 `gorm:"column:over_price" json:"over_price" form:"over_price"`
	HoldPrice      float64 `gorm:"column:hold_price" json:"hold_price" form:"hold_price"`
	Pe             float64 `gorm:"column:pe" json:"pe" form:"pe"`
	TargetPe       float64 `gorm:"column:target_pe" json:"target_pe" form:"target_pe"`
	OverPe         float64 `gorm:"column:over_pe" json:"over_pe" form:"over_pe"`
	UpdatedAt      []uint8 `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

func (StockSubscribe) TableName() string {
	return "stock_subscribe"
}

func (ss *StockSubscribe) Sprintln() string {
	res := fmt.Sprintf("[StockSubscribe: ID = %d , Symbol = %s,name=%s", ss.ID, ss.Symbol, ss.Name)
	res += "]"
	return res
}
