// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

//go:generate flex wc.l

// just like Unix wc
package main

//#include "lex.yy.h"
//#include "wc.h"
import "C"
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	flagInput = flag.String("f", "", "set input file")
)

func main() {
	flag.Parse()

	var (
		content []byte
		err     error
	)

	if *flagInput == "" {
		// cat wc.go | go run .
		content, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// go run . -f wc.go
		content, err = ioutil.ReadFile(*flagInput)
		if err != nil {
			log.Fatal(err)
		}
	}

	C.yy_scan_bytes(
		(*C.char)(C.CBytes(content)),
		C.yy_size_t(len(content)),
	)
	C.yylex()

	fmt.Printf("%8d%8d%8d\n", C.lines, C.words, C.chars)
}
