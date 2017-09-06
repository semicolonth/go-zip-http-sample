package main

import (
	"io/ioutil"
	"fmt"
	"archive/zip"
	"log"
	"bytes"
	"net/http"
	// "encoding/base64"
)

func main () {
	// Reference: https://golang.org/pkg/archive/zip/#example_Writer
	var err error

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)
	fmt.Println("buf.Len() right after initiation: ", buf.Len())
	
	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
			Name, Body string
	}{
			{"test/readme.txt", "This archive contains some text files."},
			{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
			f, err := w.Create(file.Name)
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
	err = w.Close()
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("buf.Len() after zip.Writer.Close(): ", buf.Len())
	fmt.Println("buf content: ", buf.String())

	err = ioutil.WriteFile("testwrite.zip",
		buf.Bytes(), 
		0777)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File write complete.")
	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe("localhost:4001", nil)
}

// From: https://gist.github.com/superbrothers/0a8b6390c6315916aeb8
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	// w.Header().Set("Content-Transfer-Encoding", "base64")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("testwrite.zip")
	if err != nil {
		panic(err)
	}
	// w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	// fmt.Fprint(w, base64.StdEncoding.EncodeToString(data))
	fmt.Fprintf(w, "%s", data)
}