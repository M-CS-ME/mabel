package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func srcblks(filename string) (bg, en []int) {
	file, err := os.Open(filename)
	check(err)
	in := bufio.NewScanner(file)
	var open bool = false
	for i := 0; in.Scan(); i++ {
		ln := in.Text()
		if len(ln) >= 3 {
			if ln[:3] == "```" {
				open = !open
				if open {
					bg = append(bg, i)
				} else {
					en = append(en, i)
				}
			}
		}
	}
	return
}
func write(buf []string, bg, en int) {
	for _, i := range buf[bg+1:en] {
		fmt.Println(i)
	}
}
func tangle(in string, blk int) {
	file, err := os.Open(in)
	check(err)
	f := bufio.NewScanner(file)
	bg, en := srcblks(in)
	var buf []string
	for f.Scan() {
		buf = append(buf, f.Text())
	}
	if blk == -1 {
		for i := 0; i < len(bg); i++ {
			write(buf, bg[i], en[i])
		}
	} else {
		write(buf, bg[blk], en[blk])
	}
}
func main() {
	if len(os.Args) > 2 {
		blk, _ := strconv.Atoi(os.Args[2])
		tangle(os.Args[1], blk)
	} else if len(os.Args) > 1 {
		tangle(os.Args[1], -1)
	} else {
		fmt.Println("error: no input files")
	}
}
