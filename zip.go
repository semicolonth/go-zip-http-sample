package main

import (
	"fmt"
	"archive/zip"
	"log"
	"bytes"
	"net/http"
)

func main () {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe("localhost:4001", nil)
}

// From: https://gist.github.com/superbrothers/0a8b6390c6315916aeb8
func rootHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Content-Type", "application/zip")
	w.WriteHeader(http.StatusOK)

	// Reference: https://golang.org/pkg/archive/zip/#example_Writer
	
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)
	fmt.Println("buf.Len() right after initiation: ", buf.Len())
	
	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
			Name, Body string
	}{
			{"test/readme.txt", "This archive contains some text files."},
			{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := zipWriter.Create(file.Name)
		if err != nil {
				log.Fatal(err)
		}

		var buffer []byte

		buffer = []byte(file.Body)

		// Content of buffer to be written can be read from file by ioutil.ReadFile
		// content, err := ioutil.ReadFile(file.Name)
		// if err != nil {
		// 	log.Fatal(err)
		// } else {
		// 	buffer = content
		// }

		_, err = f.Write(buffer)
		if err != nil {
				log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err = zipWriter.Close()
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("buf.Len() after zip.Writer.Close(): ", buf.Len())
	fmt.Println("buf content: ", buf.String())

	fmt.Fprintf(w, "%s", buf.String())
}