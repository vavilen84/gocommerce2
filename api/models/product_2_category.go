package models

import "api/constants"

type Product2Category struct {
	Id       int64     `json:"id" orm:"auto"`
	Product  *Product  `orm:"rel(fk)"`
	Category *Category `orm:"rel(fk)"`
}

func (m *Product2Category) TableName() string {
	return constants.Product2CategoryDBTable
}
