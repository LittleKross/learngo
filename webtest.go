package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webpage",func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w,r,"home.html")
		fmt.Fprintf(w,"Hello you've requsted: %s\n",r.URL.Path)
		log.Printf("Hello you've requsted: %s\n",r.URL.Path)
	})
	http.HandleFunc("/api/v0/",func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello you've requsted: %s\n",r.URL.Path)
		fmt.Fprintln(w,"You have requested: " + r.URL.Path)
	})
	http.ListenAndServe(":80",nil)
}
