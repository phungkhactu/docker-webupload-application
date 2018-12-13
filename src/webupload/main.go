package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func uploadFile(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		// GET
		temp, _ := template.ParseFiles("main.gtpl")

		temp.Execute(res, nil)

	} else if req.Method == "POST" {
		// Post
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(res, "%s", "Upload success") //handler.Header)
		f, err := os.OpenFile("/tmp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

	} else {
		fmt.Println("Unknown HTTP " + req.Method + "  Method")
	}
}

func main() {
	http.HandleFunc("/", uploadFile)
	http.ListenAndServe(":9090", nil) // setting listening port
}
