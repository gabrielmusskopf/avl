package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabrielmusskopf/avl/http"
	avl "github.com/gabrielmusskopf/avl/pkg"
	"github.com/gabrielmusskopf/avl/pkg/types"
)

const (
	HABILITAR_DEBUG = iota + 1
	VER_ARVORE
	INSERIR_VALOR
	INSERIR_VALORES
	BUSCAR_VALOR
	REMOVER_VALOR
	VER_POST_ORDER
	VER_PRE_ORDER
	VER_IN_ORDER
	VER_BFS
	INICIAR_HTTP
	SEE_NAME_TREE
	SEE_CPF_TREE
	SEE_BIRTH_DATE_TREE
	SEARCH_NAME
	SEARCH_CPF
	SEARCH_BIRTH_DATE
	SAIR
)

var indexDistance int
var opcoes map[int]string

func init() {
	opcoes = map[int]string{
		SAIR:                "Sair",
		HABILITAR_DEBUG:     "Habiltar debug",
		VER_ARVORE:          "Ver árvore",
		INSERIR_VALOR:       "Inserir valor",
		INSERIR_VALORES:     "Inserir valores",
		VER_POST_ORDER:      "DFS Post order",
		VER_IN_ORDER:        "DFS In order",
		VER_PRE_ORDER:       "DFS Pre order",
		VER_BFS:             "BFS",
		BUSCAR_VALOR:        "Buscar valor",
		REMOVER_VALOR:       "Remover valor",
		INICIAR_HTTP:        "Iniciar servidor HTTP",
		SEE_NAME_TREE:       "Ver árvore indexada por nome",
		SEE_CPF_TREE:        "Ver árvore indexada por CPF",
		SEE_BIRTH_DATE_TREE: "Ver árvore indexada por data de nascimento",
		SEARCH_NAME:         "Buscar por nome",
		SEARCH_CPF:          "Buscar por CPF",
		SEARCH_BIRTH_DATE:   "Buscar por data de nascimento",
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

func askValue() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(strings.TrimSuffix(input, "\n"))
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
func printPerson(p *types.Person) {
	fmt.Println()
	fmt.Printf("Nome:\t\t\t%s\n", p.Name)
	fmt.Printf("CPF:\t\t\t%s\n", p.CPF)
	fmt.Printf("RG:\t\t\t%s\n", p.RG)
	fmt.Printf("Data nascimento:\t%s\n", p.BirthDate)
	fmt.Printf("Cidade nascimento:\t%s\n", p.BirthCity)
}

func printIfExist[T avl.Ordered[T]](n *avl.IndexedTree[T, *types.Person]) {
	if n != nil {
		printPerson(n.Value)
	} else {
		fmt.Printf("\nNão existe na árvore")
	}
}

func compareWithLength(s, other types.String) int {
	if len(s) > len(other) {
		s = s[:len(other)]
	} else {
		other = other[:len(s)]
	}
	return s.Compare(other)
}

func cmdLoop(index *avl.Index) {

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
			if err := http.InitHttp(false); err != nil {
				fmt.Print(err.Error())
				break
			}
			fmt.Printf("Servidor iniciado em http://127.0.0.1:3333")
		case SEE_NAME_TREE:
			if index.Names == nil {
				fmt.Printf("Árvore vazia\n")
				continue
			}
			index.Names.PrettyPrint("")

		case SEE_CPF_TREE:
			if index.CPF == nil {
				fmt.Printf("Árvore vazia\n")
				continue
			}
			index.CPF.PrettyPrint("")

		case SEE_BIRTH_DATE_TREE:
			if index.BirthDate == nil {
				fmt.Printf("Árvore vazia\n")
				continue
			}
			index.BirthDate.PrettyPrint("")

		case SEARCH_NAME:
			fmt.Printf("Digite a chave: ")
			match := index.Names.SearchAllBy(types.String(askValue()),
				func(k1, k2 types.String) bool { return strings.HasPrefix(string(k1), string(k2)) },
				compareWithLength)

			for _, node := range match {
				printIfExist(node)
			}

		case SEARCH_CPF:
			fmt.Printf("Digite a chave: ")
			r := index.CPF.Search(types.String(askValue()))
			printIfExist(r)

		case SEARCH_BIRTH_DATE:
			fmt.Printf("Digite uma data inicial: ")
			input := askValue()
			start, err := time.Parse(types.DDMMYYYY, input)
			if err != nil {
				log.Fatalf("ERRO ao ler data %s\n", input)
			}

			fmt.Printf("Digite uma data final: ")
			input = askValue()
			end, err := time.Parse(types.DDMMYYYY, input)
			if err != nil {
				log.Fatalf("ERRO ao ler data %s\n", input)
			}

			matches := make([]*types.Person, 0)
			index.BirthDate.WalkAllBy(func(node avl.IndexedTree[types.Date, *types.Person]) {
				date, err := time.Parse(types.DDMMYYYY, types.StringFromDate(node.Key))
				if err != nil {
					log.Fatalf("ERRO no parse de %s", node.Key)
				}
				if date.After(start) && date.Before(end) {
					matches = append(matches, node.Value)
				}
			})
			if len(matches) == 0 {
				fmt.Printf("\nNão existem datas no período informado\n")
				continue
			}
			for _, node := range matches {
				printPerson(node)
			}

		case SAIR:
			fmt.Print("Desligando os motores")

		default:
			fmt.Print("Não conheço essa...")
		}
		fmt.Println()
	}
}
