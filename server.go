/*
Godoc はGoのパッケージドキュメント情報をQiitaで紹介するためのパッケージです。そのままサンプルとして使えます。
リポジトリは https://github.com/lufia/godoc-sample です。

How to read a document

ドキュメントを読むためにはgo docコマンドまたはgodocコマンドが使えます。

How to write a document

パッケージドキュメントはpackage句の直前に書く必要がありますが、
Build constraintsはpackageよりも前に書かなければなりません。
そのため、記述する順番としては、Build constraintsが先になります。

空行を入れると、別の段落として区切ることができます。

   インデントすると、ソースコードのような
   整形されたテキストも書けます

Heading

アルファベットの大文字で始まり、句読点を含まない1行だけの段落があれば、
それはヘッダとして装飾されます。ただし、ヘッダを2つ以上続けることはできません。
次の行はヘッダになりますが、その次は同じルールにも関わらず普通の段落です。

This is a header

This is not a header

*/
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ky-hy/nnecstasy-backend-go/graph"
	"github.com/ky-hy/nnecstasy-backend-go/graph/generated"
)

const defaultPort = "8080"

// It creates a GraphQL server and starts it on port 8080
func main() {
	// BUG(lufia): 保存機能は未実装です。
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}



