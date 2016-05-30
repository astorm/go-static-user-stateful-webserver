package config

import (

)

func Get(key string) string {
    var config map[string]string    
    config = make(map[string]string)
    base := "/Users/alanstorm/go/src/github.com/astorm/go-static-user-stateful-webserver"
    config["webroot"]    = base + "/static"
    config["account-db"] = base + "/accounts.db"        
    return config[key]
}