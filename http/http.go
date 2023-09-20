package http

import (
	"fmt"
	"log"
	"net"
	"net/http"
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

func index(w http.ResponseWriter, r *http.Request) {
	html := fmt.Sprintf("<ul>%s</ul>", toHtml(avl.Tree))

	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("http/tree.html"))
	err := tmpl.Execute(w, html)
	if err != nil {
		log.Fatal(err)
	}
}

func InitHttp() {
	port = "3333"
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Millisecond*100)
	if conn != nil {
		return
	}
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("http/styles"))))
	http.HandleFunc("/", index)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
