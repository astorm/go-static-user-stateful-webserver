package main

import (
    "fmt"
    _ "github.com/mattn/go-sqlite3"
	"database/sql" 
    "log"
    "os"
    "github.com/astorm/go-static-user-stateful-webserver/application"   
    "github.com/astorm/go-static-user-stateful-webserver/config"     
)

func main() {
	db, err := sql.Open("sqlite3", config.Get("account-db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    argsWithoutProg := os.Args[1:]
    
    if(len(argsWithoutProg) != 2){
        log.Fatal("USAGE: account.go username password")
    }
    var user = argsWithoutProg[0]
    var pass = argsWithoutProg[1]    
    
    var hash = application.GeneratePassword(pass)    
//    var user = "astorm"
//    var hash = application.GeneratePassword("abc123")
//    hashBytes := []byte(pass)    
//    result    := bcrypt.GenerateFromPassword(hashBytes, passBytes)


    var sql  = "INSERT INTO admin_user (username, password_hash) VALUES (?,?)"
    
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(user, hash)
	tx.Commit()
	    
    fmt.Println("Hello World")       
}