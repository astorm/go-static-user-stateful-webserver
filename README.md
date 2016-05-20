# go-static-user-stateful-webserver

This project will, eventually, be a simple static web server that serves content dynamically based on Basic HTTP Authentication.  Its specific purpose is to provide stateful composer repositories to a user and/or customer base.

I suspect this is only of interest to me, but if you've stumbled across this project, and see something that looks dumb, broken, or "wrong" from an idiomatic go perspective, issues and pull requests are welcome. 

Development Log
--------------------------------------------------
Follow along as I make really dumb beginner experiments and mistakes with go. 

- <s>Look! I reimplemented basic HTTP auth functionality that I didn't realize existed in go!</s>

- OK, using `request.BasicAuth` now.

- Not sure about this line `http.FileServer(http.Dir(folder)).ServeHTTP(responseWriter, request)` -- I'm basically doing the job of the handler function manually.  If the handler interface ever changes, I'm in trouble. 

- About to start playing around with packaging stuff.  File renames ahead.

    - Packaging -- when you import a folder it looks like go will scan the entire folder for files in that package
    - package main -- you need to `go run file1.go file2.go`
    
- Pulled in a sqllite library -- the leading `_` means the module's `init` function is called, but a `go-sqlite` symbol is not avaiable.  Avoids "derp derp you didn't use that symbol

- Pulled in a bcrypt module -- seems to use small salt, attached to pass.  Additional salt?     

- Played around with code organization -- don't include workspace stuff via ./packagename.   Still thinking of moving applicaiton.go into its own package. 

- Poor man's config() ftw