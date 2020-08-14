package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/baozhenie/study/hugo-meow/tmp.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	s,err := ioutil.ReadAll(file)
	str := string(s)
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str,":","=")
	fmt.Print(str)
}
