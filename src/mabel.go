package main

import (
	"os"
	"bufio"
	"fmt"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func tangle(in string) {
	file, err := os.Open(in)
	check(err)
	f := bufio.NewScanner(file)
	var open bool = false
	for f.Scan() {
		ln := f.Text()
		if len(ln) >= 3 {
			if ln[:3] == "```" {
				open = !open
				continue
			}
		}
		if open {
			fmt.Println(ln)
		}
	}
}
func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args[1:] {
			tangle(i)
		}
	} else {
		var file string
		fmt.Scan(&file)
		tangle(file)
	}
}
