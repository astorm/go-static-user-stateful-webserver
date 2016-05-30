package main

import (
    _ "github.com/mattn/go-sqlite3"
	"database/sql" 
	"log"
    "github.com/astorm/go-static-user-stateful-webserver/config"  	
)

func main() {
	db, err := sql.Open("sqlite3", config.Get("account-db"))
	if err != nil {
		log.Fatal(err)
	}
	
	var sql = "CREATE TABLE admin_user( "           +
        "username varchar(255) NOT NULL UNIQUE, "   +
        "password_hash varchar(60) NOT NULL)";
        
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
            
	defer db.Close()
}