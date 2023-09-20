package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gabrielmusskopf/avl"
	"github.com/gabrielmusskopf/avl/http"
)

const (
	HABILITAR_DEBUG = 1
	VER_ARVORE      = 2
	INSERIR_VALOR   = 3
	INSERIR_VALORES = 4
	BUSCAR_VALOR    = 5
	REMOVER_VALOR   = 6
	VER_POST_ORDER  = 7
	VER_PRE_ORDER   = 8
	VER_IN_ORDER    = 9
	VER_BFS         = 10
	INICIAR_HTTP    = 11
	SAIR            = 12
)

var indexDistance int
var opcoes map[int]string

func init() {
	opcoes = map[int]string{
		SAIR:            "Sair",
		HABILITAR_DEBUG: "Habiltar debug",
		VER_ARVORE:      "Ver árvore",
		INSERIR_VALOR:   "Inserir valor",
		INSERIR_VALORES: "Inserir valores",
		VER_POST_ORDER:  "DFS Post order",
		VER_IN_ORDER:    "DFS In order",
		VER_PRE_ORDER:   "DFS Pre order",
		VER_BFS:         "BFS",
		BUSCAR_VALOR:    "Buscar valor",
		REMOVER_VALOR:   "Remover valor",
		INICIAR_HTTP:    "Iniciar servidor HTTP",
	}

	avl.LogLevel = avl.NONE
	indexDistance = 5
}

func printOption(opt int, s string) {
	nsize := strconv.Itoa(opt)
	ndots := indexDistance - len(nsize)
	var dots string
	for i := 0; i < ndots; i++ {
		dots += "."
	}
	fmt.Printf("%d%s%s\n", opt, dots, s)
}

func showDebug() {
	var d string
	if avl.IsDebug() {
		d = "On"
	} else {
		d = "Off"
	}
	s := fmt.Sprintf("%s (%s)", opcoes[HABILITAR_DEBUG], d)
	printOption(HABILITAR_DEBUG, s)
}

func showMenu() {
	fmt.Println()
	for k := 1; k <= len(opcoes); k++ {
		if k == HABILITAR_DEBUG {
			showDebug()
			continue
		}
		printOption(k, opcoes[k])
	}
	fmt.Printf("\nOpção: ")
}

func askInt() int {
	var n int
	fmt.Printf("Digite o número: ")
	fmt.Scanf("%d", &n)
	return n
}

func askInts() []int {
	fmt.Print("Digite valores separados por espaço. Exemplo: 10 5 20\nDigite os números: ")
	nums := make([]int, 0)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Valor inesperado: %s. Parando por aqui...", line)
		return nums
	}

	line = line[:len(line)-1]

	for _, s := range strings.Split(line, " ") {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Valor inesperado: %s. Parando por aqui...", s)
			break
		}
		nums = append(nums, i)
	}
	return nums
}

func doOr(n *avl.TreeNode, fail, sucess string) {
	if n != nil {
		fmt.Print(sucess)
	} else {
		fmt.Print(fail)
	}
}

func main() {
	fmt.Printf("Árvore AVL\n")

	opt := -1

	for opt != SAIR {
		showMenu()
		fmt.Scanf("%d", &opt)

		switch opt {
		case HABILITAR_DEBUG:
			avl.ToggleDebug()

		case VER_ARVORE:
			if avl.Tree == nil {
				fmt.Printf("Árvore vazia\n")
				continue
			}
			avl.Tree.PrettyPrint("")

		case INSERIR_VALOR:
			d := askInt()
			avl.Tree = avl.Tree.Add(d)
			doOr(avl.Tree, "Não pôde adicionar", "Ok!")

		case INSERIR_VALORES:
			ds := askInts()
			if avl.Tree == nil {
				avl.Tree = avl.FromArray(ds)
			} else {
                avl.Tree = avl.Tree.AddFromArray(ds)
            }
			doOr(avl.Tree, "Não pôde criar árvore", "Ok!")

		case BUSCAR_VALOR:
			d := askInt()
			doOr(avl.Tree.Serach(d), "Não existe na árvore", "Existe na árvore")

		case REMOVER_VALOR:
			d := askInt()
			avl.Tree = avl.Tree.Remove(d)
			fmt.Print("Ok!")

		case VER_PRE_ORDER:
			avl.Tree.PreOrder()

		case VER_IN_ORDER:
			avl.Tree.InOrder()

		case VER_POST_ORDER:
			avl.Tree.PostOrder()

		case VER_BFS:
			avl.Tree.BFS()

		case INICIAR_HTTP:
			go http.InitHttp()
			fmt.Printf("Servidor iniciado em http://127.0.0.1:3333")

		case SAIR:
			fmt.Print("Desligando os motores")

		default:
			fmt.Print("Não conheço essa...")
		}
		fmt.Println()
	}
}

//Ideia: servidor HTTP para visualização da árvore
