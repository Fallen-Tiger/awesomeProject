// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+
func main() {
	var str string
	var begin1 = time.Now()
	for index, a := range os.Args {
		fmt.Println(index)
		str += string(index) + "    " + a + "\n"
	}
	var end1 = time.Now()
	fmt.Println(str)
	fmt.Println(end1.Sub(begin1).Milliseconds())
	str = ""
	var begin2 = time.Now()
	for index, a := range os.Args {
		str += string(index) + "    " + a + "\n"
	}
	var end2 = time.Now()
	fmt.Println(str)
	fmt.Println(end2.Sub(begin2).Milliseconds())
}

//!-
