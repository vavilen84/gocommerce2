package models

import (
	"api/constants"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"time"
)

type Product struct {
	BaseModel
	Title string `json:"title" orm:"column(title)"`
	SKU   string `json:"sku" orm:"column(sku)"`
	Price int    `json:"price" orm:"column(price)"`

	Categories []*Category `orm:"rel(m2m);rel_through(api/models.PostTags)"`
}

func (m *Product) TableName() string {
	return constants.ProductDBTable
}

func (m *Product) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateSKUAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.ProductModel)
		return false
	}
	return true
}

func (m *Product) validateCommonFields(valid *validation.Validation) {
	valid.MaxSize(m.Title, 255, "title")
	valid.Required(m.Title, "title")

	valid.MaxSize(m.SKU, 255, "title")
	valid.Required(m.SKU, "sku")

	valid.Match(m.SKU, regexp.MustCompile(`^[a-z0-9_-]*$`), "sku")

	valid.Required(m.Price, "price")

	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.UpdatedAt, "updated_at")
}

func (m *Product) ValidateSKUAlreadyUsed(o orm.Ormer, valid *validation.Validation) {
	productFromDb, err := FindProductBySKU(o, m.SKU)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (productFromDb.Id != 0) && (m.Id != productFromDb.Id) {
			err := valid.SetError("sku", "SKU is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *Product) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateProductExists(o, &valid)
	m.ValidateSKUAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.ProductModel)
		return false
	}
	return true
}

func (m *Product) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func InsertProduct(o orm.Ormer, m *Product) (err error) {
	m.clearValidationErrors()
	m.setTimestampsOnCreate()
	isValid := m.validateOnInsert(o)
	if !isValid {
		return
	}
	_, err = o.Insert(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindProductById(o orm.Ormer, id int64) (m Product, err error) {
	err = o.QueryTable(constants.ProductDBTable).Filter("id", id).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindProductBySKU(o orm.Ormer, sku string) (m Product, err error) {
	err = o.QueryTable(constants.ProductDBTable).Filter("sku", sku).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *Product) ValidateProductExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindProductById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("Product with id #%d does not exist", m.Id)
			valid.SetError("Product", errMsg)
		}
	}
}

func UpdateProduct(o orm.Ormer, m *Product) (err error) {
	m.clearValidationErrors()
	m.setTimestampsOnUpdate()
	isValid := m.validateOnUpdate(o)
	if !isValid {
		return
	}
	_, err = o.Update(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func DeleteProduct(o orm.Ormer, m *Product) error {
	_, err := FindProductById(o, m.Id)
	if err != nil {
		logs.Error(err)
		return err
	}
	_, err = o.Delete(m)
	if err != nil {
		logs.Error(err)
	}
	return nil
}
