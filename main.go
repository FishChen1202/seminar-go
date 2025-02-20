package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 設定首頁路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World! Test Verstion 3")
	})

	// 啟動伺服器，監聽在 8080 埠
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
