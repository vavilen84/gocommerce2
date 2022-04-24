package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddProductTable_20220424_132114 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddProductTable_20220424_132114{}
	m.Created = "20220424_132114"

	migration.Register("AddProductTable_20220424_132114", m)
}

// Run the migrations
func (m *AddProductTable_20220424_132114) Up() {
	m.SQL(`
		CREATE TABLE product
		(
			id         INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
			title      varchar(255) NOT NULL,
			sku        varchar(255) NOT NULL,
			price      BIGINT UNSIGNED NOT NULL,
			created_at INT(11) NOT NULL,
			updated_at INT(11) NOT NULL
		) ENGINE=InnoDB CHARSET=utf8;
	`)
}

// Reverse the migrations
func (m *AddProductTable_20220424_132114) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
