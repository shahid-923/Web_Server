**Web_Server – Go HTTP Server**

A lightweight and beginner-friendly HTTP server built in Go, demonstrating how to:

**Serve static files**
1.Handle GET and POST routes
2.Read form input
3.Use Go’s net/http package to build simple web applications
4.This project is ideal for beginners learning backend development with Go.

**Project Structure**
Web_Server/
│── main.go
│
└── static/
    │── index.html
    │── form.html

**Features**
✔ Serve Static Files
The server automatically serves files inside the static/ folder (HTML, CSS, JS).

✔ Custom Routes
/hello → Handles GET requests
/form → Handles POST form submissions
/ → Serves index.html

✔ Form Handling
Accepts user input (name & address) and responds dynamically.

✔ Simple, clean Go code
Uses only Go’s in-built standard library — no external packages.

**Installation & Usage**
1. Clone the repository
git clone https://github.com/shahid-923/Web_Server.git

2. Navigate to the project
cd Web_Server

3. Run the server
go run main.go

Your server starts at:

👉 http://localhost:8080/
**🔗 Endpoints Overview**
**Route	Method	Description**
/	GET	Serves index.html
/hello	GET	Returns “Hello!”
/form	POST	Accepts name and address fields


**Code Summary**
**Serve static files**
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/", fileServer)

Handle /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w, "Hello!")
}

Handle /form
func formHandler(w http.ResponseWriter, r *http.Request){
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

**📝 HTML Files**
index.html
A simple welcome page.
form.html
Form that sends POST data:

<form method="POST" action="/form">
    <label>Name</label><input type="text" name="name"/>
    <label>Address</label><input type="text" name="address"/>
    <input type="submit" value="Submit"/>
</form>

**📦Technologies Used**
Go (Golang)
HTML
net/http (Go Standard Library)

**🤝Contributing**
Contributions and suggestions are welcome.
Feel free to fork the repo and submit a pull request.

**📜License**
This project is open-source and available under the MIT License.
