package models

import (
	"api/constants"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_ = iota
	UserRoleCustomer
	UserRoleAdmin
)

type User struct {
	Id        int64  `json:"id" orm:"auto"`
	Email     string `json:"email" orm:"column(email);unique"`
	Password  string `json:"password" orm:"column(password)"`
	Salt      string `json:"salt" orm:"column(salt)"`
	Role      int    `json:"role" orm:"column(role)"`
	FirstName string `json:"first_name" orm:"column(first_name)"`
	LastName  string `json:"last_name" orm:"column(last_name)"`

	CreatedAt int  `json:"created_at" orm:"column(created_at)"`
	UpdatedAt int  `json:"updated_at" orm:"column(updated_at)"`
	DeletedAt *int `json:"deleted_at" orm:"column(deleted_at)"`
}

func (m *User) FindByEmail(o orm.Ormer) (err error) {
	err = o.QueryTable(constants.UserModel).Filter("email", m.Email).One(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindUserById(o orm.Ormer, id int64) (m User, err error) {
	err = o.QueryTable(constants.UserModel).Filter("id", id).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func InsertUser(o orm.Ormer, m *User) (err error) {
	m.setPassword()
	m.setTimestampsOnCreate()
	isValid := m.validateOnInsert(o)
	if !isValid {
		return
	}
	id, err := o.Insert(m)
	if err != nil {
		logs.Error(err)
		return
	}
	m.Id = id
	return
}

func UpdateUser(o orm.Ormer, m User) (err error) {
	m.setPasswordOnUpdate(o)
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

func (m *User) setPasswordOnUpdate(o orm.Ormer) {
	oldUser, err := FindUserById(o, m.Id)
	if err != nil {
		logs.Error(err)
	}
	if m.Password != "" {
		m.setPassword()
	} else {
		m.Password = oldUser.Password
		m.Salt = oldUser.Salt
	}
}

func (m *User) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *User) setTimestampsOnUpdate() {
	now := int(time.Now().Unix())
	m.UpdatedAt = now
}

func (m *User) setPassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.Salt = salt
}

func (m *User) validateEmailAlreadyInUse(o orm.Ormer, valid *validation.Validation) {
	u := User{Email: m.Email}
	err := u.FindByEmail(o)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (u.Id != 0) && (u.Id != m.Id) {
			err := valid.SetError("email", "Email is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *User) ValidateUserExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindUserById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("User with id #%d does not exist", m.Id)
			valid.SetError("User", errMsg)
		}
	}
}

func (m *User) validateCommonFields(valid *validation.Validation) {
	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")
	valid.Email(m.Email, "email")

	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")

	valid.Required(m.Salt, "salt")

	valid.Required(m.Role, "salt")
	valid.Range(m.Role, UserRoleCustomer, UserRoleAdmin, "role")

	valid.Required(m.FirstName, "first_name")
	valid.MaxSize(m.FirstName, 255, "first_name")

	valid.Required(m.LastName, "last_name")
	valid.MaxSize(m.LastName, 255, "last_name")
}

func (m *User) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	valid.Required(m.Id, "id")
	m.ValidateUserExists(o, &valid)
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err)
		}
		return false
	}
	return true
}

func (m *User) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err)
		}
		return false
	}
	return true
}
