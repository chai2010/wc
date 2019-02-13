/* Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved. */
/* Use of this source code is governed by a Apache */
/* license that can be found in the LICENSE file. */

/* just like Unix wc */

%option noyywrap

%{
#include "lex.h"
#include "wc.h"

#include <string.h>
%}

%%

[a-zA-Z]+ { words++; chars += strlen(yytext); }
\n        { chars++; lines++; }
.         { chars++; }

%%

void wc_yy_scan_bytes(const void* p, int len) {
	yy_scan_bytes(p, len);
}

int wc_yylex(void) {
	return yylex();
}
