# Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

run: lex.yy.c lex.yy.h
	@go fmt
	@go vet

	go run . -f wc.go
	cat wc.go | go run .

lex.yy.c lex.yy.h: wc.l
	flex --prefix=yy --header-file=lex.yy.h wc.l

clean:
	-rm lex.yy.c
