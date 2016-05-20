package main

//TODO: create database in code
//TODO: command/function to add user account/password
//    CREATE TABLE admin_user(
//    username varchar(255) NOT NULL,
//    password_hash varchar(60) NOT NULL);
//TODO: unique index (sqlite equivalent?) on username in table
//TODO: additional salt?
//TODO: move application to package?
import (
    //"fmt"
)

func main() {
    taskWebServer()
}

func config(key string) string {
    var config map[string]string    
    config = make(map[string]string)
    base := "/Users/alanstorm/go/src/github.com/astorm/go-static-user-stateful-webserver"
    config["webroot"]    = base + "/static"
    config["account-db"] = base + "/accounts.db"        
    return config[key]
}