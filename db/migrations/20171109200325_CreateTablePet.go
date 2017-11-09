
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20171109200325(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS `pet` ("+
	  	"`id` int(11) NOT NULL AUTO_INCREMENT,"+
	  	"`name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,"+
	  	"`age` int(5) NOT NULL,"+
	  	"`photo` text COLLATE utf8_unicode_ci NOT NULL,"+
	  	"`updated_at` datetime DEFAULT NULL,"+
	  	"`created_at` datetime DEFAULT NULL,"+
		"PRIMARY KEY (`id`)"+
	") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;")
	if err != nil {
		return err
	}
	return nil
}

// Down is executed when this migration is rolled back
func Down_20171109200325(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE `pet`;")
	if err != nil {
		return err
	}
	return nil
}
