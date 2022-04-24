package models

import (
	"api/constants"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"time"
)

type Product struct {
	BaseModel
	Title string `json:"title" orm:"column(title)"`
	SKU   string `json:"sku" orm:"column(sku)"`
	Price int    `json:"price" orm:"column(price)"`
}

func (m *Product) TableName() string {
	return constants.ProductDBTable
}

func (m *Product) validateOnInsert() error {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.ProductModel)
	}
	return nil
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

func (m *Product) validateOnUpdate(o orm.Ormer) error {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err)
		}
		e := errors.New(fmt.Sprintf("Model %v is not valid", constants.ProductModel))
		return e
	}
	return nil
}

func (m *Product) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func InsertProduct(o orm.Ormer, m *Product) error {
	m.clearValidationErrors()
	m.setTimestampsOnCreate()
	err := m.validateOnInsert()
	if err != nil {
		logs.Error(err)
		return err
	}
	_, err = o.Insert(m)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
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

func (m *Product) Delete(o orm.Ormer) error {
	err := m.validateProductExists(o)
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
