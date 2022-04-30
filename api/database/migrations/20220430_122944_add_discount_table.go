package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddDiscountTable_20220430_122944 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddDiscountTable_20220430_122944{}
	m.Created = "20220430_122944"

	migration.Register("AddDiscountTable_20220430_122944", m)
}

// Run the migrations
func (m *AddDiscountTable_20220430_122944) Up() {
	m.SQL("CREATE TABLE `discount` (" +
		"id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"title varchar(255) NOT NULL, " +
		"type INT NOT NULL, " +
		"amount INT DEFAULT NULL, " +
		"percentage INT DEFAULT NULL, " +
		"created_at INT(11)," +
		"updated_at INT(11)" +
		") ENGINE=InnoDB CHARSET=utf8;")
}

// Reverse the migrations
func (m *AddDiscountTable_20220430_122944) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
