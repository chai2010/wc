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

最重要的词法分析函数的实现由flex工具生成:

```c
extern int yylex(void);
```

创建`wc.l`文件:

```flex
%option noyywrap

%{
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
$ flex --prefix=yy --header-file=lex.yy.h wc.l
```

然后在Go语言中调用词法分析器并输出结果:

```go
package main

//#include "lex.yy.h"
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

## 改进: 从内存读取数据

flex生成的代码导出了`yyin`和`yyout`全局变量，可以用于指定打开的输入流文件（否则从标准输入读取）：

```c
extern FILE *yyin, *yyout;
``

不过以上的是C语言文件接口，在Go语言中使用比较麻烦。此外flex还生成了以下三个函数：

```c
YY_BUFFER_STATE yy_scan_buffer (char *base,yy_size_t size  );
YY_BUFFER_STATE yy_scan_string (yyconst char *yy_str  );
YY_BUFFER_STATE yy_scan_bytes (yyconst char *bytes,yy_size_t len  );
```

然后调整Go语言函数：

```go
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
```

在调用`C.yylex`函数之前调用`C.yy_scan_bytes`函数用Go语言的切片初始化词法扫描的缓冲区。
