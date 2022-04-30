package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddTaxTable_20220430_142805 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddTaxTable_20220430_142805{}
	m.Created = "20220430_142805"

	migration.Register("AddTaxTable_20220430_142805", m)
}

// Run the migrations
func (m *AddTaxTable_20220430_142805) Up() {
	m.SQL("CREATE TABLE `tax` (" +
		"id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"title varchar(255) NOT NULL, " +
		"amount INT DEFAULT NULL, " +
		"percentage INT DEFAULT NULL, " +
		"created_at INT(11)," +
		"updated_at INT(11)" +
		") ENGINE=InnoDB CHARSET=utf8;")
}

// Reverse the migrations
func (m *AddTaxTable_20220430_142805) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
