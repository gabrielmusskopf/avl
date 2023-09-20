package main

import (
	"flag"
	"fmt"

	"github.com/gabrielmusskopf/avl"
	"github.com/gabrielmusskopf/avl/http"
)

func main() {
	avl.TreeEvents = &avl.Queue[string]{}

	justHttp := flag.Bool("http", false, "just serve http, no command line")
	flag.Parse()

	if *justHttp {
		fmt.Println("Servidor iniciado em http://127.0.0.1:3333")
		if err := http.InitHttp(true); err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println("Terminando servidor")
		return
	}

	fmt.Printf("Árvore AVL\n")
	cmdLoop()
}
