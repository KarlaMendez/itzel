func login(res http.ResponseWriter, req *http.Request) {
    // If method is GET serve an html login page
    if req.Method != "POST" {
        http.ServeFile(res, req, "index.html")
        return
    }    
}
