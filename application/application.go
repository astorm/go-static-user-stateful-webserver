package application

import (
	"io"
	"net/http"
	"fmt"
    _ "github.com/mattn/go-sqlite3"
	"database/sql" 
	"log"   
	"golang.org/x/crypto/bcrypt"
    "github.com/astorm/go-static-user-stateful-webserver/config"
	
)

func parseUsernameAndPasswordFromAuthHeader(request *http.Request) map[string]string {
    var usernameAndPassword map[string]string
    usernameAndPassword = make(map[string]string)
    usernameAndPassword["username"] = ""
    usernameAndPassword["password"] = ""    
    username, password, ok := request.BasicAuth()
    if(ok == false){
        return usernameAndPassword
    }    
    usernameAndPassword["username"] = username
    usernameAndPassword["password"] = password    
    return usernameAndPassword;
}

func debugRequest(request *http.Request) {
    fmt.Printf("%+v\n", request.URL);
}

func sendAuthRequiredHeaders(responseWriter http.ResponseWriter) {
    responseWriter.Header().Set("WWW-Authenticate","Basic realm=\"Composer\"")
    responseWriter.WriteHeader(http.StatusUnauthorized)
}

func writeBodyBodyUnauthorizedRequest(responseWriter http.ResponseWriter) {
	io.WriteString(responseWriter, "Authorization Required!")	
}

func authenticateUsenameAndPassword(usernameAndPassword map[string]string) bool {
	db, err := sql.Open("sqlite3", config.Get("account-db"))
	fmt.Println(config.Get("account-db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	stmt, err := db.Prepare(
	    "SELECT password_hash FROM admin_user WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(usernameAndPassword["username"])
	if err != nil {
		log.Fatal(err)
	}
		
	defer rows.Close()
	for rows.Next() {
	    var password_hash string;
		err = rows.Scan(&password_hash)
		if err != nil {
			log.Fatal(err)
		}    
        passBytes := []byte(usernameAndPassword["password"])
        hashBytes := []byte(password_hash)    
        result    := bcrypt.CompareHashAndPassword(hashBytes, passBytes)
        return result == nil //nil means the password checked out
	}
    return false
}

func processRequest(responseWriter http.ResponseWriter, request *http.Request) {

    usernameAndPassword := parseUsernameAndPasswordFromAuthHeader(request);
    if(usernameAndPassword["username"] == "" || usernameAndPassword["password"] == ""){
        sendAuthRequiredHeaders(responseWriter)
        writeBodyBodyUnauthorizedRequest(responseWriter)
        return;
    }             
    
    if(authenticateUsenameAndPassword(usernameAndPassword)){
        //var folder = "./static/" + usernameAndPassword["username"]
        var folder = config.Get("webroot") + "/" + usernameAndPassword["username"]
        fmt.Println(folder)
        http.FileServer(http.Dir(folder)).ServeHTTP(responseWriter, request)
        //io.WriteString(responseWriter, "Autehnticated!")
    } else{
        sendAuthRequiredHeaders(responseWriter)
        io.WriteString(responseWriter, "Invalid Username/Password")
    }
        
    debugRequest(request);
}

func processLoginRequest(responseWriter http.ResponseWriter, request *http.Request) {

    sendAuthRequiredHeaders(responseWriter)
        writeBodyBodyUnauthorizedRequest(responseWriter)

    //sendAuthRequiredHeaders(responseWriter)
    io.WriteString(responseWriter, "Implement Me")
}

func GeneratePassword(pass string) []byte {
    //pass        := "abc123"
    passBytes   := []byte(pass)
    var result, err = bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
    if err != nil {
      log.Fatal(err)
    }  
    return result
}

func TaskWebServer() {
   http.HandleFunc("/", processRequest)
   http.HandleFunc("/login", processLoginRequest)
   http.ListenAndServe(":8000", nil)	
}

func taskHelloGoodbye() {
    fmt.Printf("Hello %s\n", "world");	    
    fmt.Printf("Goodbye %s\n", "world");
}