package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func ErrorIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/error")
}

func FileServerHandler(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			ErrorIndex(w, r)
			return
		}
		if r.URL.Path[1:] == "" {
			_, _ = fmt.Fprintf(w, "<h1>Hi, there!</h1>")
			fsh.ServeHTTP(w, r)
		} else {
			fsh.ServeHTTP(w, r)
		}
	})
}

func main() {
	http.Handle("/", FileServerHandler(http.Dir("./")))

	fmt.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
