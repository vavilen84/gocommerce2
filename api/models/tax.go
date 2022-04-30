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

type Tax struct {
	BaseModel
	Title      string `json:"title" column:"title"`
	Amount     uint32 `json:"amount" column:"amount"`
	Percentage int    `json:"percentage" column:"percentage"`
}

func (m *Tax) GetTableName() string {
	return constants.TaxDBTable
}

func (m *Tax) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.TaxModel)
		return false
	}
	return true
}

func (m *Tax) validateCommonFields(valid *validation.Validation) {
	valid.MaxSize(m.Title, 255, "title")
	valid.Required(m.Title, "title")

	if m.Amount == 0 && m.Percentage == 0 {
		valid.SetError("amount", "Amount or Percentage should be set")
		valid.SetError("percentage", "Amount or Percentage should be set")
	}

	valid.Max(m.Percentage, 100, "percentage")

	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.UpdatedAt, "updated_at")
}

func (m *Tax) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateTaxExists(o, &valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.TaxModel)
		return false
	}
	return true
}

func (m *Tax) ValidateTitleAlreadyUsed(o orm.Ormer, valid *validation.Validation) {
	taxFromDb, err := FindTaxByTitle(o, m.Title)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (taxFromDb.Id != 0) && (m.Id != taxFromDb.Id) {
			err := valid.SetError("title", "Title is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *Tax) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func InsertTax(o orm.Ormer, m *Tax) (err error) {
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

func FindTaxByTitle(o orm.Ormer, title string) (m Tax, err error) {
	err = o.QueryTable(constants.TaxDBTable).Filter("title", title).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindTaxById(o orm.Ormer, id int64) (m Tax, err error) {
	err = o.QueryTable(constants.TaxDBTable).
		Filter("id", id).
		One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *Tax) ValidateTaxExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindTaxById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("Tax with id #%d does not exist", m.Id)
			valid.SetError("Tax", errMsg)
		}
	}
}

func UpdateTax(o orm.Ormer, m *Tax) (err error) {
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

func DeleteTax(o orm.Ormer, m *Tax) error {
	_, err := FindTaxById(o, m.Id)
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
