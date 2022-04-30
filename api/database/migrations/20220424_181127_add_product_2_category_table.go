package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddProduct2CategoryTable_20220424_181127 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddProduct2CategoryTable_20220424_181127{}
	m.Created = "20220424_181127"

	migration.Register("AddProduct2CategoryTable_20220424_181127", m)
}

// Run the migrations
func (m *AddProduct2CategoryTable_20220424_181127) Up() {
	m.SQL("CREATE TABLE `product_2_category` (" +
		"id INT NOT NULL PRIMARY KEY AUTO_INCREMENT, " +
		"product_id INT NOT NULL, " +
		"category_id INT NOT NULL, " +
		"created_at INT(11)," +
		"updated_at INT(11)," +
		"UNIQUE KEY product_category_id (product_id,category_id)" +
		") ENGINE=InnoDB CHARSET=utf8;")

	m.SQL("ALTER TABLE product_2_category " +
		"ADD CONSTRAINT fk_product_id FOREIGN KEY (product_id) " +
		"REFERENCES product(id) ON DELETE CASCADE;")

	m.SQL("ALTER TABLE product_2_category " +
		"ADD CONSTRAINT fk_category_id FOREIGN KEY (category_id) " +
		"REFERENCES category(id) ON DELETE CASCADE;")
}

// Reverse the migrations
func (m *AddProduct2CategoryTable_20220424_181127) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
