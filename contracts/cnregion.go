package contracts

import "github.com/libtnb/region"

type Region interface {
	// ParseByCode 通过编码解析出省市区街道
	ParseByCode(code string) (province, city, area, street string, err error)
	// ParseByName 通过省市区街道反向解析出编码
	ParseByName(province, city, area string, street ...string) (code string, err error)
	// Search 通过关键字搜索省市区街道编码
	Search(keyword string) (result []region.Region)
}
