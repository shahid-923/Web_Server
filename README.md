**Web_Server — Simple HTTP Server in Go**

A lightweight Go-based web server that demonstrates handling routes, serving static files, and processing HTML form submissions.
This project is ideal for beginners learning backend fundamentals in Go using the net/http package.

🚀 **Features**
Serves static HTML files from a static/ directory
Handles multiple routes (/, /hello, /form)
Processes form input using POST requests
Uses Go’s built-in net/http package — no external dependencies

*Clean and simple project structure for easy learning*

📁 **Project Structure**
Web_Server/
│-- main.go
└───static/
      │-- index.html
      │-- form.html

🔗 **Available Endpoints**
Endpoint	Method	Description
/	GET	Serves the static home page
/hello	GET	Returns a simple greeting message
/form	POST	Accepts form data (name, address) and displays it

🧠 **Concepts You Will Learn**

How to create routes in Go
How to serve static files (HTML, CSS, JS)
How to handle form submissions
How Go processes GET vs POST requests
How HTTP handlers work internally

▶️ **How to Run**

Clone the repository
Navigate into the project folder
Run the server:
go run main.go
Open in browser:
http://localhost:8080/

📦 **Tech Used**

Go (Golang)
net/http package
HTML for frontend pages

📝 **Future Improvements**

Add CSS for styling
Add more routes
Add database support
Create APIs (JSON responses)
