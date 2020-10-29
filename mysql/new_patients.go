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
    query := `CREATE TABLE patients(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50),
    password VARCHAR(120),
	patientname VARCHAR(120)
   );`
    if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }
	
	
	
}