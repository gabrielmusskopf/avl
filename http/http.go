package http

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var port string
var routes *Routes

func InitHttp() error {
	port = "3333"
	conn, _ := net.DialTimeout("tcp", net.JoinHostPort("", port), time.Millisecond*100)
	if conn != nil {
		return fmt.Errorf("Porta %s indisponível. Servidor ou outro processo em execução", port)
	}
	if routes != nil {
		routes = &Routes{}
	}

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("http/styles"))))
	http.HandleFunc("/api/avl/delete", routes.handleDelete)
	http.HandleFunc("/api/avl/add", routes.handleAdd)
	http.HandleFunc("/api/avl/clear", routes.handleClear)
	http.HandleFunc("/", routes.handleTree)

	go func() {
        _ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
    }()
    return nil
}
