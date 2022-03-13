package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddUserTable_20220312_182028 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddUserTable_20220312_182028{}
	m.Created = "20220312_182028"

	migration.Register("AddUserTable_20220312_182028", m)
}

// Run the migrations
func (m *AddUserTable_20220312_182028) Up() {
	m.SQL("CREATE TABLE `user` (" +
		"id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"email VARCHAR(255), " +
		"password TEXT, " +
		"salt TEXT," +
		"first_name VARCHAR(255)," +
		"last_name VARCHAR(255)," +
		"role SMALLINT," +
		"created_at INT(11)," +
		"updated_at INT(11)," +
		"deleted_at INT(11)" +
		");")
	m.SQL("ALTER TABLE user ADD UNIQUE INDEX email_idx(email);")
}

// Reverse the migrations
func (m *AddUserTable_20220312_182028) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
