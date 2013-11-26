// chat room example
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fzzy/sockjs-go/sockjs"
)

func chatHandler(s sockjs.Session) {

	client := login(s)
	if err := clients.Add(client); err != nil {
		client.Send("Username is taken.")
		return
	}
	defer clients.Remove(client)
	client.Send(fmt.Sprintf("Welcome, %s.", client.Name))

	for {
		m := s.Receive()
		if m == nil {
			break
		}
		languages.Broadcast(client, m)
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
