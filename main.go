package main

import(
	"fmt"
	"net/http"
	"log"
)
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	} 
	if r.Method != "GET" {
	http.Error(w, "Method is not supported.", http.StatusNotFound)    // dont allow user to modify hello 
	return
    }    
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
    fmt.Printf("POST request successful\n") 
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address) 
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))	// handles serving static files from the "static" directory
	http.Handle("/", fileServer)	                        // route all requests to the file server
	http.HandleFunc("/form", formHandler)                  // handle form submissions at "/form" endpoint
	http.HandleFunc("/hello", helloHandler)                // handle requests at "/hello" endpoint
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	} 
}