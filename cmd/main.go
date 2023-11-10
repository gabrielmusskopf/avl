package main

import (
	"flag"
	"fmt"

	"github.com/gabrielmusskopf/avl/http"
	avl "github.com/gabrielmusskopf/avl/pkg"
	"github.com/gabrielmusskopf/avl/pkg/data"
)

const N_DATA int = 1000

func main() {
	avl.TreeEvents = &avl.Queue[string]{}
	reader := &data.CsvPersonReader{}
	if !data.HasData() {
        fmt.Println("Gerando dados...")
		data.Generate(N_DATA)
        fmt.Println("Dado gerados!")
	}
	people := reader.Read(data.DataPath)
	index := avl.BuildIndexes(people)

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
	cmdLoop(index)
}
