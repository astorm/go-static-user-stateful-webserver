package main

import (
	"io"
	"net/http"
	"fmt"
//	"strings"
//	"encoding/base64"
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
    responseWriter.Header().Set("WWW-Authenticate","Basic realm=\"My Realm\"")
    responseWriter.WriteHeader(http.StatusUnauthorized)
}

func writeBodyBodyUnauthorizedRequest(responseWriter http.ResponseWriter) {
	io.WriteString(responseWriter, "Authorization Required!")	
}

func authenticateUsenameAndPassword(usernameAndPassword map[string]string) bool {
    if(usernameAndPassword["username"] == "astorm" && usernameAndPassword["password"] == "pass") {
        return true
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
        var folder = "./static/" + usernameAndPassword["username"]
        http.FileServer(http.Dir(folder)).ServeHTTP(responseWriter, request)
        //io.WriteString(responseWriter, "Autehnticated!")
    } else{
        sendAuthRequiredHeaders(responseWriter)
        io.WriteString(responseWriter, "Invalid Username/Password")
    }
        
    debugRequest(request);
}

func processRequest2(responseWriter http.ResponseWriter, request *http.Request) {
    //http.FileServer(http.Dir("./static")).ServeHTTP(responseWriter, request)
}
 
func main() {
	//http.HandleFunc("/", processRequest2)
	
	http.HandleFunc("/", processRequest)
	http.ListenAndServe(":8000", nil)
	
    //http.Handle("/", http.FileServer(http.Dir("./static")))
    //http.ListenAndServe(":3000", nil)	
	
}

//    http.Handle(
//        "/static/", 
//        http.StripPrefix(
//            "/static/", 
//            http.FileServer(
//                http.Dir("./public")
//            )
//        )
//    );
