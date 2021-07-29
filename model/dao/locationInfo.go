package dao

import (
	xdb "ImyouboHomeKit/utils/db"
	"time"
)

const LocationInfoTable = "location_info"

type LocationInfo struct {
	Id            uint64    `json:"id" xorm:"id"`
	Name          string    `json:"name" xorm:"name"`
	EnName        string    `json:"enName" xorm:"en_name"`
	CountryCode   string    `json:"countryCode" xorm:"country_code"`
	CountryName   string    `json:"countryName" xorm:"country_name"`
	CountryEnName string    `json:"countryEnName" xorm:"country_en_name"`
	Adm1Name      string    `json:"adm1Name" xorm:"adm1_name"`
	Adm1EnName    string    `json:"adm1EnName" xorm:"adm1_en_name"`
	Adm2Name      string    `json:"adm2Name" xorm:"adm2_name"`
	Adm2EnName    string    `json:"adm2EnName" xorm:"adm2_en_name"`
	Latitude      string    `json:"latitude" xorm:"latitude"`
	Longitude     string    `json:"longitude" xorm:"longitude"`
	AdCode        string    `json:"adCode" xorm:"ad_code"`
	Ct            time.Time `json:"ct" xorm:"ct"`
	Ut            time.Time `json:"ut" xorm:"ut"`
}

func GetMultiLocationInfo(db *xdb.DB, where string, data ...interface{}) ([]*LocationInfo, error) {
	var lis []*LocationInfo
	defer db.Close()
	return lis, db.Table(LocationInfoTable).Where(where, data...).Find(&lis)
}

func GetOneLocationInfo(db *xdb.DB, where string, data ...interface{}) (*LocationInfo, error) {
	location := new(LocationInfo)

	if has, err := db.Table(LocationInfoTable).Where(where, data...).Get(location); err != nil {
		return nil, err
	} else if !has {
		return nil, xdb.EmptyResultErr
	}
	defer db.Close()
	return location, nil
}
