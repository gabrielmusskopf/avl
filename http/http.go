package http

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gabrielmusskopf/avl"
)

var port string

/*
<div>
    <ul>
        <li> elemento </li>
        <ul>
            <li> elemento </li>
            <li> elemento </li>
        </ul>
    </ul>
</div>
*/

func toHtml(n *avl.TreeNode) string {
	if n == nil {
		return ""
	}

	builder := strings.Builder{}

	if n.Left != nil || n.Right != nil {
		builder.WriteString("<ul>")
	}
	if n.Left != nil {
		builder.WriteString(toHtml(n.Left))
	}
	if n.Right != nil {
		builder.WriteString(toHtml(n.Right))
	}
	if n.Left != nil || n.Right != nil {
		builder.WriteString("</ul>")
	}

	return fmt.Sprintf(`<li><div>%d</div>%s</li>`, n.Value, builder.String())
}

func handleTree(w http.ResponseWriter, r *http.Request) {
	html := fmt.Sprintf("<ul>%s</ul>", toHtml(avl.Tree))

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("http/tree.html"))
	err := tmpl.Execute(w, html)
	if err != nil {
		log.Fatal(err)
	}
}

func parseNode(r *http.Request) (int, error) {
    node := r.URL.Query().Get("node")
    if node == "" {
        return 0, fmt.Errorf("parâmetro 'node' não encontrado")
    }
    num, err := strconv.Atoi(node)
    if err != nil {
        return 0, fmt.Errorf("nó '%s' inválido", node)
    }
    return num, nil
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
    node, err := parseNode(r)
    if err != nil {
        avl.Error(err.Error())
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    avl.Tree = avl.Tree.Remove(node)
    handleTree(w, r)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
    node, err := parseNode(r)
    if err != nil {
        avl.Error(fmt.Sprintf("%s\n", err.Error()))
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    avl.Tree = avl.Tree.Add(node)
    handleTree(w, r)
}

func handleClear(w http.ResponseWriter, r *http.Request) {
    avl.Tree = nil
    handleTree(w, r)
}

func InitHttp() {
	port = "3333"
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Millisecond*100)
	if conn != nil {
		return
	}
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("http/styles"))))
	http.HandleFunc("/api/avl/delete", handleDelete)
	http.HandleFunc("/api/avl/add", handleAdd)
	http.HandleFunc("/api/avl/clear", handleClear)
	http.HandleFunc("/index", handleTree)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
