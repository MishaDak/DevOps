package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ReadTextFromFile(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive request path: %s\n", r.URL.Path)
	if r.URL.Path[1:] == "" {
		files, err := os.ReadDir("/tmp/files/")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Fprintf(w, "<h1> Hi, there!</h1>")
			htmlstr := "<a href=" + file.Name() + ">" + file.Name() + "</a>\n"
			fmt.Fprintf(w, htmlstr)
		}
	} else {
		sff := "/tmp/files/" + r.URL.Path
		file, err := os.Open(sff)

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "not file")
		} else {
			defer file.Close()

			data := make([]byte, 64)

			for {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				fmt.Fprintf(w, string(data[:n]))
			}
		}
	}
}
func main() {
	fmt.Printf("Server started")
	http.HandleFunc("/", ReadTextFromFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
