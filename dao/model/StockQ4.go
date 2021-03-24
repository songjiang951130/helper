package model

type StockQ4 struct {
	ID         uint32  `gorm:"primaryKey "`
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
