package models

type Location struct {
	Ip string `json:"ip"`
	ContinentCode string `json:"continent_code"`
	ContinentName string `json:"continent_name"`
	CountryCode2 string `json:"country_code2"`
	CountryCode3 string `json:"country_code3"`
	CountryName string `json:"country_name"`
	CountryCapital string `json:"country_capital"`
	StateProvince string `json:"state_prov"`
	City string `json:"city"`
	Zipcode string `json:"zipcode"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Isp string `json:"isp"`
	Organization string `json:"organization"`
}
