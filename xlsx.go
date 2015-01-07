package main

import (
	"archive/zip"
	"fmt"
	// "strings"
	"io"
	"log"
	"os"
)

func main() {

	// Open a zip archive for reading.
	r, err := zip.OpenReader("test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		if len(f.Name) >= 15 {

			if f.Name[:14] == "xl/worksheets/" {
				fmt.Printf("Contents of %s:\n", f.Name)
				rc, err := f.Open()
				if err != nil {
					log.Fatal(err)
				}
				_, err = io.CopyN(os.Stdout, rc, 1000)
				if err != nil {
					log.Fatal(err)
				}
				rc.Close()
				fmt.Println("\n")
			}
		}
	}

}
