package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gabrielmusskopf/avl"
)

type Routes struct {
}

/*
How each node should be:
<div>
  <ul>
    <li>node</li>
    <ul>
      <li>node</li>
      <li>//another node//</li>
    </ul>
  </ul>
</div>
*/

func nodeToHtml(n *avl.TreeNode, children string) string {
	if n == nil {
		return `<li><div class="node"><p></p><div></div></div></li>`
	}
	return fmt.Sprintf(`<li><div class="node"><p>%d</p><div>%d</div></div>%s</li>`, n.BF, n.Value, children)
}

func toHtml(n *avl.TreeNode) string {
	if n == nil {
		return nodeToHtml(n, "")
	}

	children := strings.Builder{}
	if n.Left != nil || n.Right != nil {
		children.WriteString("<ul>")
		children.WriteString(toHtml(n.Left))
		children.WriteString(toHtml(n.Right))
		children.WriteString("</ul>")
	}

	return nodeToHtml(n, children.String())
}

type HtmlTree struct {
	Tree   string
	Events []string
}

func (ro *Routes) handleTree(w http.ResponseWriter, r *http.Request) {
	tree := fmt.Sprintf("<ul>%s</ul>", toHtml(avl.Tree))
	events := make([]string, 0)
	avl.TreeEvents.Walk(func(t string) {
		events = append(events, t)
	})
	htmlTree := &HtmlTree{
		Tree:   tree,
		Events: events,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("http/tree.html"))
	err := tmpl.Execute(w, htmlTree)
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

func (ro *Routes) handleDelete(w http.ResponseWriter, r *http.Request) {
	node, err := parseNode(r)
	if err != nil {
		avl.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	avl.Tree = avl.Tree.Remove(node)
	ro.handleTree(w, r)
}

func (ro *Routes) handleAdd(w http.ResponseWriter, r *http.Request) {
	node, err := parseNode(r)
	if err != nil {
		avl.Error(fmt.Sprintf("%s\n", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	avl.Tree = avl.Tree.Add(node)
	ro.handleTree(w, r)
}

func (ro *Routes) handleClear(w http.ResponseWriter, r *http.Request) {
	avl.Tree = nil
	ro.handleTree(w, r)
}
