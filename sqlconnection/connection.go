package sqlconnection

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"




var db *sql.DB
var err error



// sql connection 
func DbConn() (db *sql.DB) {
   db, err :=sql.Open("mysql", "root:@(127.0.0.1:3306)/test?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
    return db
}


