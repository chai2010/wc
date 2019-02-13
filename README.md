# 基于flex实现的wc命令, 单词计数

创建`wc.h`文件, 包含要统计的变量和`yylex`词法解析函数:

```c
extern int chars;
extern int words;
extern int lines;
```

创建`wc.c`文件, 保护变量的定义并初始化为0值:

```c
#include "wc.h"

int chars = 0;
int words = 0;
int lines = 0;
```

最重要的是`yylex`函数在`lex.h`文件声明, 函数的实现由flex工具生成:

```c
extern int yylex(void);
```

创建`wc.l`文件:

```flex
%option noyywrap

%{
#include "lex.h"
#include "wc.h"
%}

%%

[a-zA-Z]+ { words++; chars += strlen(yytext); }
\n        { chars++; lines++; }
.         { chars++; }

%%
```

通过以下命令生成`lex.yy.c`文件(其中包含`yylex`函数的实现):

```
$ flex wc.l
```

然后在Go语言中调用词法分析器并输出结果:

```go
package main

//#include "lex.h"
//#include "wc.h"
import "C"
import "fmt"

func main() {
	C.yylex()
	fmt.Printf("%8d%8d%8d\n", C.lines, C.words, C.chars)
}
```

通过以下的命令运行:

```
$ cat wc.l | go run .
      19      47     355
```
