package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddCategoryTable_20220424_173424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddCategoryTable_20220424_173424{}
	m.Created = "20220424_173424"

	migration.Register("AddCategoryTable_20220424_173424", m)
}

// Run the migrations
func (m *AddCategoryTable_20220424_173424) Up() {
	m.SQL(`
		CREATE TABLE category
		(
			id         INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title      varchar(255) NOT NULL,
			created_at INT(11) NOT NULL,
			updated_at INT(11) NOT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`)
}

// Reverse the migrations
func (m *AddCategoryTable_20220424_173424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
