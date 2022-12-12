---
title: "Ginのmiddlewareの呼び出し順序"
date: 2022-12-12T21:20:37+09:00
draft: false
comments: false
adsense: false
---

GoのWebフレームワークであるGinについて勘違いしていたので備忘録としてブログに残す。

middlewareってLIFOのイメージがあったのでGinも当然そうかと思っていたが違っていた。

```go
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.WriteString("2")
		ctx.Next()
		ctx.Writer.WriteString("4")
	})
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.WriteString("1")
		ctx.Next()
		ctx.Writer.WriteString("5")
	})
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.Writer.WriteString("3")
	})
	if err := router.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

上記のコードでサーバーを起動し、`localhost:8080/hello`にアクセスした際にレスポンスとして`12345`が返ることを期待したが実際は下記の通り。

```sh
% curl localhost:8080/hello
21354%
```

どうやら`Use`で呼び出すごとに引数の関数がappendされていき、実行時にはindexをインクリメントしながら呼び出しているよう。
FIFOだったのである。

なので`12345`と出力したいのであれば下記のように実装する。

```go
package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.WriteString("1")
		ctx.Next()
		ctx.Writer.WriteString("5")
	})
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.WriteString("2")
		ctx.Next()
		ctx.Writer.WriteString("4")
	})
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.Writer.WriteString("3")
	})
	if err := router.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

```sh
% curl localhost:8080/hello
12345%
```
