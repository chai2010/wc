// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate flex wc.l

// just like Unix wc
package main

//#include "wc.h"
import "C"
import "fmt"

func main() {
	C.yylex()
	fmt.Printf("%8d%8d%8d\n", C.lines, C.words, C.chars)
}
