# Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

run: lex.yy.c
	cat wc.l | go run .

lex.yy.c:
	flex wc.l

clean:
	-rm lex.yy.c
