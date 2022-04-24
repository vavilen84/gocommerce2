package models

type Product2Category struct {
	BaseModel
	Product  *Product  `orm:"rel(fk)"`
	Category *Category `orm:"rel(fk)"`
}
