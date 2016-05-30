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
   "github.com/astorm/go-static-user-stateful-webserver/application"
)

func main() {
 //   type ConfigFunction func(string)
    application.TaskWebServer()
}