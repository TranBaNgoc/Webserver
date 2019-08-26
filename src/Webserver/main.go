// Server 1 is minimal "echo" server 
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // mỗi request tới sẽ gọi tới hàm handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// in ra url cua request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}