// chat room example
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fzzy/sockjs-go/sockjs"
)

var users *sockjs.SessionPool = sockjs.NewSessionPool()

func chatHandler(s sockjs.Session) {
	users.Add(s)
	defer users.Remove(s)

	client := login(s)
	defer delete(clients, client.Name)
	for {
		m := s.Receive()
		if m == nil {
			break
		}
		m = []byte(fmt.Sprintf("%s: %s", client.Name, m))
		users.Broadcast(m)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func Start() {
	mux := sockjs.NewServeMux(http.DefaultServeMux)
	conf := sockjs.NewConfig()
	http.Handle("/static", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", indexHandler)
	mux.Handle("/chat", chatHandler, conf)

	log.Println("The server is up an running at http://0.0.0.0:8081")
	err := http.ListenAndServe("0.0.0.0:8081", mux)
	if err != nil {
		fmt.Println(err)
	}
}
