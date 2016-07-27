package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {

	findStr := flag.String("find", "q264", "string to find")
	replStr := flag.String("replace", "H264", "string to replace found string, must be same length > 0")
	prefix := flag.String("prefix", "M", "prefix to new file")
	flag.Parse()
	files := flag.Args()

	// check flags and arguments
	new_files, ok := validArgs(*findStr, *replStr, files)
	if !ok {
		log.Fatal("ERROR with arguments.")
	}

	for _, fname := range new_files {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Printf("\nERROR reading: %s\n", fname)
		}

		replaceString(*findStr, *replStr, &data)
		err = ioutil.WriteFile(*prefix+fname, data, 0644)
		if err != nil {
			log.Printf("\nERROR writing: %s\n", *prefix+fname)
		}

		log.Printf("Processed file: %s\n", fname)
	}

}

func validArgs(find, replace string, files []string) (new_files []string, ok bool) {
	fmt.Printf("find: %s, replace: %s, files: %v\n", find, replace, files)

	if len(find) != len(replace) {
		log.Printf("strings %s %s are not equal sized\n", find, replace)
		return files, false
	}

	if len(find) == 0 {
		log.Printf("strings are zero sized\n")
		return files, false
	}

	if len(files) == 0 {
		log.Printf("No files to process\n")
		return files, false
	}

	new_files = make([]string, 0)
	for _, s := range files {
		if strings.Contains(s, "*") {
			more_files, err := filepath.Glob(s)
			if err != nil {
				log.Printf("error expanding a wildcard %s\n", s)
			}
			new_files = append(new_files, more_files...)
		} else { // regular file name
			new_files = append(new_files, s)
		}
	}

	return new_files, true
}

func replaceString(find, replace string, data *[]byte) {
	f := []byte(find)
	r := []byte(replace)
	lf := len(find)
	d := *data
	ld := len(d)

	for i := 0; i < ld-lf+1; i++ {
		if d[i] == f[0] {
			fw := d[i:(i + lf)]
			if reflect.DeepEqual(f, fw) {
				for j, b := range r {
					d[i+j] = b
				}
			}
		}
	}
}
