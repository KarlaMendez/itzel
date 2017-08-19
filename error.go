func login(res http.ResponseWriter, req *http.Request) {
    // If method is GET serve an html login page
    if req.Method != "POST" {
        http.ServeFile(res, req, "login.html")
        return
    }    

    // Grab the username/password from the submitted post form
    username := req.FormValue("username")
    password := req.FormValue("password")

    // Grab from the database 
    var databaseUsername  string
    var databasePassword  string

    // Search the database for the username provided
    // If it exists grab the password for validation
    err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
    // If not then redirect to the login page
    if err != nil {
        http.Redirect(res, req, "/login", 301)
        return
    }
    
    // Validate the password
    err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
    // If wrong password redirect to the login
    if err != nil {
        http.Redirect(res, req, "/login", 301)
        return
    }

    // If the login succeeded
    res.Write([]byte("Hello " + databaseUsername))
}