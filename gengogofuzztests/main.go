package main

import (
	gogopop "github.com/gogo/fuzztests/gengogofuzztests/gogopop"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	path := os.Args[1]
	randoms := gogopop.Randoms()
	for i, r := range randoms {
		if err := ioutil.WriteFile(filepath.Join(path, strconv.Itoa(i)), r, 0666); err != nil {
			panic(err)
		}
	}
}
