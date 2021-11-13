package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webpage",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"home.html")
		fmt.Fprintf(w,"Hello you've requsted: %s\n",r.URL.Path)
		log.Printf("Hello you've requsted: %s\n",r.URL.Path)
	})
	r.HandleFunc("/api/v0/path/current",func(w http.ResponseWriter, r *http.Request) {
		log.Printf("The current path is: %s\n",r.URL.Path)
		fmt.Fprintln(w,"The current path is: " + r.URL.Path)
	})
	http.ListenAndServe(":80",r)
}
