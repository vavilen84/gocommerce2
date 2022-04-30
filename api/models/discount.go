package models

import (
	"api/constants"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Discount struct {
	BaseModel
	Title      string `json:"title" column:"title"`
	Amount     uint32 `json:"amount" column:"amount"`
	Percentage int    `json:"percentage" column:"percentage"`
	Type       int    `json:"type" column:"type"`
}

func GetDiscountTypesList() []int {
	return []int{
		constants.DiscountCartType,
		constants.DiscountCategoryType,
		constants.DiscountProductType,
	}
}

func (m *Discount) GetTableName() string {
	return constants.DiscountDBTable
}

func (m *Discount) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.DiscountModel)
		return false
	}
	return true
}

func (m *Discount) validateCommonFields(valid *validation.Validation) {
	valid.MaxSize(m.Title, 255, "title")
	valid.Required(m.Title, "title")

	valid.Required(m.Type, "type")
	if m.Type != constants.DiscountCartType &&
		m.Type != constants.DiscountCategoryType &&
		m.Type != constants.DiscountProductType {
		valid.SetError("type", "Type is not valid")
	}

	if m.Amount == 0 && m.Percentage == 0 {
		valid.SetError("amount", "Amount or Percentage should be set")
		valid.SetError("percentage", "Amount or Percentage should be set")
	}

	valid.Max(m.Percentage, 100, "percentage")

	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.UpdatedAt, "updated_at")
}

func (m *Discount) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateDiscountExists(o, &valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.DiscountModel)
		return false
	}
	return true
}

func (m *Discount) ValidateTitleAlreadyUsed(o orm.Ormer, valid *validation.Validation) {
	discountFromDb, err := FindDiscountByTitle(o, m.Title)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (discountFromDb.Id != 0) && (m.Id != discountFromDb.Id) {
			err := valid.SetError("title", "Title is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *Discount) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func InsertDiscount(o orm.Ormer, m *Discount) (err error) {
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

func FindDiscountByTitle(o orm.Ormer, title string) (m Discount, err error) {
	err = o.QueryTable(constants.DiscountDBTable).Filter("title", title).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindDiscountById(o orm.Ormer, id int64) (m Discount, err error) {
	err = o.QueryTable(constants.DiscountDBTable).
		Filter("id", id).
		One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *Discount) ValidateDiscountExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindDiscountById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("Discount with id #%d does not exist", m.Id)
			valid.SetError("Discount", errMsg)
		}
	}
}

func UpdateDiscount(o orm.Ormer, m *Discount) (err error) {
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

func DeleteDiscount(o orm.Ormer, m *Discount) error {
	_, err := FindDiscountById(o, m.Id)
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
