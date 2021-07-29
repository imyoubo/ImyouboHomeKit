package model

import (
	"ImyouboHomeKit/api"
	"ImyouboHomeKit/model/dao"
	xdb "ImyouboHomeKit/utils/db"
	"fmt"
	"strings"
)

func GetAllLocationInfo() ([]*dao.LocationInfo, error) {
	db, err := xdb.GetDefaultDB()
	if err != nil {
		return nil, err
	}
	return dao.GetMultiLocationInfo(db, "")
}

func GetLocationById(locationId uint64) (*dao.LocationInfo, error) {
	db, err := xdb.GetDefaultDB()
	if err != nil {
		return nil, err
	}
	return dao.GetOneLocationInfo(db, "id = ?", locationId)
}

func ListLocationInfo(param *api.ListLocationInfoRequest)  ([]*dao.LocationInfo, int64, error)  {
	db, err := xdb.GetDefaultDB()
	if err != nil {
		return nil, 0, err
	}
	where := strings.Builder{}
	where.WriteString(" 1 = 1 ")
	if param.Kw != "" {
		where.WriteString(fmt.Sprintf("and ( name like '%%%s%%' or lower(en_name) like lower('%%%s%%')) ", param.Kw, param.Kw))
	}
	total, err := db.Table(dao.LocationInfoTable).Where(where.String()).Count()
	if err != nil {
		return nil, 0, err
	}
	var l []*dao.LocationInfo
	return l, total, db.Table(dao.LocationInfoTable).Where(where.String()).Limit(param.Limit, param.Offset).Find(&l)
}
