package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddJwtInfoTable_20220422_174443 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddJwtInfoTable_20220422_174443{}
	m.Created = "20220422_174443"

	migration.Register("AddJwtInfoTable_20220422_174443", m)
}

// Run the migrations
func (m *AddJwtInfoTable_20220422_174443) Up() {
	m.SQL("CREATE TABLE `jwt_info` (" +
		"id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"email VARCHAR(255), " +
		"password TEXT, " +
		"salt TEXT," +
		"user_id VARCHAR(255)," +
		"last_name VARCHAR(255)," +
		"role SMALLINT," +
		"created_at INT(11)," +
		"updated_at INT(11)" +
		");")
	m.SQL("ALTER TABLE user ADD UNIQUE INDEX email_idx(email);")
}

// Reverse the migrations
func (m *AddJwtInfoTable_20220422_174443) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
