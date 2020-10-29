package main

import (
    "database/sql"
 
    "log"
 

    _ "github.com/go-sql-driver/mysql"
)

func main() {            
//sql.Open("mysql", "sql7372682:VYI9QS4jk2@tcp(sql7.freemysqlhosting.net:3306)/sql7372682?parseTime=true")
    db, err :=sql.Open("mysql", "root:@(127.0.0.1:3306)/test?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    { // Create a new table
		query :=`CREATE TABLE appointment (
		id int(6) unsigned NOT NULL AUTO_INCREMENT,
		docName varchar(30) NOT NULL,
		available bool DEFAULT true,
		dateTime TIMESTAMP ,
		patient varchar(30),
		PRIMARY KEY (id)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`
		if _, err := db.Exec(query); err != nil {
				log.Fatal(err)
		}
    }
	
}